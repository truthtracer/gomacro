/*
 * gomacro - A Go interpreter with Lisp-like macros
 *
 * Copyright (C) 2017-2019 Massimiliano Ghilardi
 *
 *     This Source Code Form is subject to the terms of the Mozilla Public
 *     License, v. 2.0. If a copy of the MPL was not distributed with this
 *     file, You can obtain one at http://mozilla.org/MPL/2.0/.
 *
 *
 * lookup.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	r "reflect"

	"github.com/truthtracer/gomacro/go/etoken"
	"github.com/truthtracer/gomacro/go/types"
	"github.com/truthtracer/gomacro/go/typeutil"
)

type depthMap struct {
	gmap typeutil.Map
}

func (m *depthMap) visited(gtype types.Type, depth int) bool {
	if at := m.gmap.At(gtype); at != nil && at.(int) < depth {
		// already visited at shallower depth.
		// avoids infinite loop for self-referencing types
		// as type X struct { *X }
		return true
	}
	m.gmap.Set(gtype, depth)
	return false
}

// FieldByName returns the (possibly embedded) struct field with given name,
// and the number of fields found at the same (shallowest) depth: 0 if not found.
// Private fields are returned only if they were declared in pkgpath.
func (t *xtype) FieldByName(name, pkgpath string) (field StructField, count int) {
	if name == "_" || t.kind != r.Struct {
		return
	}
	// debugf("field cache for %v <%v> = %v", unsafe.Pointer(t), t, t.cache.field)
	qname := QName2(name, pkgpath)

	v := t.universe
	if v.ThreadSafe {
		defer un(lock(v))
	}
	field, found := t.cache.field[qname]
	if found {
		if field.Index == nil { // marker for ambiguous field names
			count = int(field.Offset) // reuse Offset as "number of ambiguous fields"
		} else {
			count = 1
		}
		return field, count
	}
	var tovisit []StructField
	var visited depthMap
	field, count, tovisit = fieldByName(t, qname, 0, nil, &visited)

	// breadth-first recursion
	for count == 0 && len(tovisit) != 0 {
		var next []StructField
		for _, f := range tovisit {
			efield, ecount, etovisit := fieldByName(unwrap(f.Type), qname, f.Offset, f.Index, &visited)
			if count == 0 {
				if ecount > 0 {
					field = efield
				} else {
					// no recursion if we found something
					next = append(next, etovisit...)
				}
			}
			count += ecount
		}
		tovisit = next
	}
	if count > 0 {
		cacheFieldByName(t, qname, &field, count)
	}
	return field, count
}

func fieldByName(t *xtype, qname QName, offset uintptr, index []int, m *depthMap) (field StructField, count int, tovisit []StructField) {
	// also support embedded fields: they can be named types or pointers to named types
	t, gtype := derefStruct(t)
	if gtype == nil || m.visited(gtype, len(index)) {
		return
	}
	// debugf("fieldByName: visiting %v <%v> <%v> at depth %d", t.kind, t.gtype, t.rtype, len(index))

	n := t.NumField()
	for i := 0; i < n; i++ {

		gfield := gtype.Field(i)
		if matchFieldByName(qname, gfield) {
			if count == 0 {
				field = t.field(i) // lock already held. makes a copy
				field.Offset += offset
				field.Index = concat(index, field.Index) // make a copy of index
				// debugf("fieldByName: %d-th field of <%v> matches: %#v", i, t.rtype, field)
			}
			count++
		} else if count == 0 && gfield.Anonymous() {
			efield := t.field(i) // lock already held
			efield.Offset += offset
			efield.Index = concat(index, efield.Index) // make a copy of index
			// debugf("fieldByName: %d-th field of <%v> is anonymous: %#v", i, t.rtype, efield)
			tovisit = append(tovisit, efield)
		}
	}
	return field, count, tovisit
}

func derefStruct(t *xtype) (*xtype, *types.Struct) {
	if t != nil {
		switch gtype := t.gtype.Underlying().(type) {
		case *types.Struct:
			return t, gtype
		case *types.Pointer:
			gelem, ok := gtype.Elem().Underlying().(*types.Struct)
			if ok {
				// not t.Elem(), it would acquire Universe lock
				return unwrap(t.elem()), gelem
			}
		}
	}
	return nil, nil
}

// return true if gfield name matches given name, or if it's anonymous and its *type* name matches given name
func matchFieldByName(qname QName, gfield *types.Var) bool {
	// always check the field's package, not the type's package
	if qname == QNameGo(gfield) {
		return true
	}
	if gfield.Anonymous() {
		gtype := gfield.Type()
		if gptr, ok := gtype.(*types.Pointer); ok {
			// unnamed field has unnamed pointer type, as for example *Foo
			// check the element type
			gtype = gptr.Elem()
		}
		switch gtype := gtype.(type) {
		case *types.Basic:
			// is it possible to embed basic types?
			// yes, and they work as unexported embedded fields,
			// i.e. in the same package as the struct that includes them
			return qname == QNameGo2(gtype.Name(), gfield.Pkg())
		case *types.Named:
			// gtype.Obj().Pkg() and gfield.Pkg() should be identical for *unexported* fields
			// (they are ignored for exported fields)
			return qname == QNameGo2(gtype.Obj().Name(), gfield.Pkg())
		}
	}
	return false
}

// add field to type's cache.field. used by Type.FieldByName after a successful lookup
func cacheFieldByName(t *xtype, qname QName, field *StructField, count int) {
	if t.cache.field == nil {
		t.cache.field = make(map[QName]StructField)
	}
	if count > 1 {
		field.Index = nil             // marker for ambiguous field names
		field.Offset = uintptr(count) // reuse Offset as "number of ambiguous fields"
	}
	t.cache.field[qname] = *field
	t.universe.cache.field = true
}

// anonymousFields returns the anonymous fields of a struct type (either named or unnamed)
// also accepts a pointer to a struct type
func anonymousFields(t *xtype, offset uintptr, index []int, m *depthMap) []StructField {
	t, gtype := derefStruct(t)
	if gtype == nil || m.visited(gtype, len(index)) {
		return nil
	}
	n := gtype.NumFields()
	var tovisit []StructField
	for i := 0; i < n; i++ {
		gfield := gtype.Field(i)
		if gfield.Anonymous() {
			field := t.field(i) // not t.Field(), it would acquire Universe lock
			field.Offset += offset
			field.Index = concat(index, field.Index) // make a copy of index
			tovisit = append(tovisit, field)
		}
	}
	return tovisit
}

// MethodByName returns the method with given name (including wrapper methods for embedded fields)
// and the number of methods found at the same (shallowest) depth: 0 if not found.
// Private methods are returned only if they were declared in pkgpath.
func (t *xtype) MethodByName(name, pkgpath string) (method Method, count int) {
	// debugf("method cache for %v <%v> = %v", unsafe.Pointer(t), t, t.cache.method)

	v := t.universe
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return t.methodByName(name, pkgpath)
}

// If generics v2 are disabled, only interfaces, named types, structs, pointers to named types, and pointers to structs can have methods.
// If generics v2 are enabled, they add a few methods to most types => every type can have methods.
//
// Note: unnamed structs can have types because they can embed anonymous fields with methods.
//
// canHaveMethods() can only be called while t.universe is locked
func canHaveMethods(t *xtype) bool {
	if t == nil {
		return false
	}
	if etoken.GENERICS.V2_CTI() || t.kind == r.Interface {
		return true
	}
	if t.kind == r.Ptr {
		t = unwrap(t.elem())
	}
	return t != nil && t.kind == r.Struct || t.Named()
}

func (t *xtype) methodByName(name, pkgpath string) (method Method, count int) {
	if name == "_" || !canHaveMethods(t) {
		return
	}
	qname := QName2(name, pkgpath)
	method, found := t.cache.method[qname]
	// debugf("methodByName type = %v, qname = %v: found = %v method in cache = %v", t, qname, found, method)
	if found {
		index := method.Index
		if index < 0 { // marker for ambiguous method names
			count = -index
		} else {
			count = 1
		}
		return method, count
	}
	var visited depthMap
	method, count = methodByName(t, qname, nil)
	if count == 0 {
		tovisit := anonymousFields(t, 0, nil, &visited)
		// breadth-first recursion on struct's anonymous fields
		for count == 0 && len(tovisit) != 0 {
			var next []StructField
			for _, f := range tovisit {
				et := unwrap(f.Type)
				emethod, ecount := methodByName(et, qname, f.Index)
				if count == 0 {
					if ecount > 0 {
						method = emethod
					} else {
						// no recursion if we found something
						next = append(next, anonymousFields(et, f.Offset, f.Index, &visited)...)
					}
				}
				count += ecount
			}
			tovisit = next
		}
	}
	if count > 0 {
		cacheMethodByName(t, qname, &method, count)
	}
	return method, count
}

// For interfaces, search in *all* methods including wrapper methods for embedded interfaces
// For all other named types, only search in explicitly declared methods, ignoring wrapper methods for embedded fields.
func methodByName(t *xtype, qname QName, index []int) (method Method, count int) {

	// debugf("methodByName: visiting %v <%v> <%v> at depth %d", t.kind, t.gtype, t.rtype, len(index))

	// also support embedded fields: they can be interfaces, named types, pointers to named types
	if t.kind == r.Ptr {
		te := unwrap(t.elem())
		if te.kind == r.Interface || te.kind == r.Ptr {
			return
		}
		t = te
	}
	n := t.NumMethod()
	// fmt.Printf("looking up method %v in type %v\n", qname, t)
	for i := 0; i < n; i++ {
		gmethod := t.gmethod(i)
		if matchMethodByName(qname, gmethod) {
			// fmt.Printf("  match   : method %v\n", QNameGo(gmethod))
			if count == 0 {
				method = t.method(i)                                 // lock already held
				method.FieldIndex = concat(index, method.FieldIndex) // make a copy of index
				// debugf("methodByName: %d-th explicit method of <%v> matches: %#v", i, t.rtype, method)
			}
			count++
		} else {
			// fmt.Printf("  mismatch: method %v\n", QNameGo(gmethod))
		}
	}
	return
}

// return true if gmethod name matches given name
func matchMethodByName(qname QName, gmethod *types.Func) bool {
	// always check the methods's package, not the type's package
	return qname == QNameGo(gmethod)
}

// add method to type's cache.method. used by Type.MethodByName after a successful lookup
func cacheMethodByName(t *xtype, qname QName, method *Method, count int) {
	if t.cache.method == nil {
		t.cache.method = make(map[QName]Method)
	}
	if count > 1 {
		method.Index = -count // marker for ambiguous method names
	}
	t.cache.method[qname] = *method
	t.universe.cache.method = true
}

// visit type's direct and embedded fields in breadth-first order
func (v *Universe) VisitFields(t Type, visitor func(StructField)) {
	xt := unwrap(t)
	if xt == nil {
		return
	}
	var curr, tovisit []*xtype
	curr = []*xtype{xt}
	var seen typeutil.Map

	for len(curr) != 0 {
		for _, xt := range curr {
			// embedded fields can be named types or pointers to named types
			xt, _ = derefStruct(xt)
			if xt == nil || xt.kind != r.Struct || seen.At(xt.gtype) != nil {
				continue
			}
			seen.Set(xt.gtype, xt.gtype)

			for i, n := 0, xt.NumField(); i < n; i++ {
				field := xt.field(i)
				visitor(field)
				if field.Anonymous {
					tovisit = append(tovisit, unwrap(field.Type))
				}
			}
		}
		curr = tovisit
		tovisit = nil
	}
}

func invalidateCache(gtype types.Type, t interface{}) {
	if t, ok := t.(Type); ok {
		t := unwrap(t)
		t.cache.field = nil
		t.cache.method = nil
	}
}

func invalidateMethodCache(gtype types.Type, t interface{}) {
	if t, ok := t.(Type); ok {
		t := unwrap(t)
		t.cache.method = nil
	}
}

// clears all xtype.cache.field and xtype.cache.method.
// invoked by NamedOf() when a type is redefined.
func (v *Universe) InvalidateCache() {
	if v.cache.field || v.cache.method {
		v.gmap.Iterate(invalidateCache)
		v.cache.field = false
		v.cache.method = false
	}
}

// clears all xtype.cache.method.
// invoked by AddMethod() when a method is redefined.
func (v *Universe) InvalidateMethodCache() {
	if v.cache.method {
		v.gmap.Iterate(invalidateMethodCache)
		v.cache.method = false
	}
}
