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
 * named.go
 *
 *  Created on May 07, 2017
 *      Author Massimiliano Ghilardi
 */

package xreflect

import (
	"go/token"
	r "reflect"

	"github.com/truthtracer/gomacro/go/etoken"

	"github.com/truthtracer/gomacro/go/types"
)

// NamedOf returns a new named type for the given type name and package.
// Initially, the underlying type may be set to interface{} - use SetUnderlying to change it.
// These two steps are separate to allow creating self-referencing types,
// as for example type List struct { Elem int; Rest *List }
func (v *Universe) NamedOf(name, pkgpath string) Type {
	if v.ThreadSafe {
		defer un(lock(v))
	}
	return v.namedOf(name, pkgpath)
}

func (v *Universe) namedOf(name, pkgpath string) Type {
	return v.reflectNamedOf(name, pkgpath, rTypeOfForward)
}

// alternate version of namedOf(), to be used when reflect.Type is known
func (v *Universe) reflectNamedOf(name, pkgpath string, rtype r.Type) Type {
	opt := OptDefault
	underlying := v.BasicTypes[rtype.Kind()]
	if underlying == nil {
		underlying = v.TypeOfForward
		opt = OptIncomplete
	}
	// debugf("namedof: %s/%s, rtype = %v, %v", pkgpath, name, rtype, opt)
	pkg := v.loadPackage(pkgpath)
	typename := types.NewTypeName(token.NoPos, (*types.Package)(pkg), name, nil)
	return v.maketype4(
		// kind is reflect.Invalid;
		// underlying.GoType() will often be inaccurate and equal to interface{};
		// rtype will often be inaccurate and equal to TypeOfForward.
		// All these issues will be fixed by Type.SetUnderlying()
		r.Invalid,
		// if etoken.GENERICS_V2_CTI, v.BasicTypes[kind] is a named type
		// wrapping the actual basic type
		types.NewNamed(typename, underlying.GoType().Underlying(), nil),
		rtype,
		opt,
	)
}

// SetUnderlying sets the underlying type of a named type and marks t as complete.
// It panics if the type is unnamed, or if the underlying type is named,
// or if SetUnderlying() was already invoked on the named type.
func (t *xtype) SetUnderlying(underlying Type) {
	gtype, ok := t.gtype.(*types.Named)
	if !ok {
		xerrorf(t, "SetUnderlying of unnamed type %v", t)
	}
	v := t.universe
	if t.kind != r.Invalid || gtype.Underlying() != v.TypeOfForward.GoType() || t.rtype != v.TypeOfForward.ReflectType() {
		// redefined type. try really hard to support it.
		v.InvalidateCache()
		// xerrorf(t, "SetUnderlying invoked multiple times on named type %v", t)
	}
	xunderlying := unwrap(underlying)
	gunderlying := xunderlying.gtype.Underlying() // in case underlying is named
	t.kind = gtypeToKind(xunderlying, gunderlying)
	gtype.SetUnderlying(gunderlying)
	// debugf("SetUnderlying: updated <%v> reflect Type from <%v> to <%v> (%v)", gtype, t.rtype, underlying.ReflectType(), t.option)
	t.rtype = underlying.ReflectType()
	if t.kind == r.Interface {
		// propagate methodvalue from underlying interface to named type
		t.methodvalue = xunderlying.methodvalue
		t.cache.method = nil
		t.cache.field = nil
	} else if etoken.GENERICS.V2_CTI() {
		v.addTypeMethodsCTI(t)
	}
	if t.option == OptIncomplete {
		t.option = OptDefault
	}
}

// AddMethod adds method 'name' to type.
// It panics if the type is unnamed, or if the signature is not a function type,
// Returns the method index, or < 0 in case of errors
func (t *xtype) AddMethod(name string, signature Type) int {
	// debugf("AddMethod on type %v, method %s %v", t, name, signature)
	gtype, ok := t.gtype.(*types.Named)
	if !ok {
		xerrorf(t, "AddMethod on unnamed type %v", t)
	}
	kind := gtypeToKind(t, gtype.Underlying())
	if kind == r.Ptr || kind == r.Interface {
		xerrorf(t, "AddMethod: cannot add methods to named %s type: <%v>", kind, t)
	}
	if signature.Kind() != r.Func {
		xerrorf(t, "AddMethod on <%v> of non-function: %v", t, signature)
	}
	gsig := signature.gunderlying().(*types.Signature)
	// accept both signatures "non-nil receiver" and "nil receiver, use the first parameter as receiver"
	grecv := gsig.Recv()
	if grecv == nil && gsig.Params().Len() != 0 {
		grecv = gsig.Params().At(0)
	}
	if grecv == nil {
		xerrorf(t, "AddMethod on <%v> of function with no receiver and no parameters: %v", t, gsig)
	}
	if !types.IdenticalIgnoreTags(grecv.Type(), gtype) &&
		// !types.IdenticalIgnoreTags(grecv.Type(), gtype.Underlying()) &&
		!types.IdenticalIgnoreTags(grecv.Type(), types.NewPointer(gtype)) {

		label := "receiver"
		if gsig.Recv() == nil {
			label = "first parameter"
		}
		xerrorf(t, "AddMethod on <%v> of function <%v> with mismatched %s type: %v", t, gsig, label, grecv.Type())
	}

	gpkg := gtype.Obj().Pkg()
	gfun := types.NewFunc(token.NoPos, gpkg, name, gsig)

	n1 := gtype.NumMethods()
	index := gtype.ReplaceMethod(gfun)
	n2 := gtype.NumMethods()

	for len(t.methodvalue) < n2 {
		t.methodvalue = append(t.methodvalue, r.Value{})
	}
	// store in t.methodvalue[index] a nil function with the correct reflect.Type:
	// needed by Type.GetMethod(int) to retrieve the method's reflect.Type
	//
	// fixes gophernotes issue 174
	t.methodvalue[index] = r.Zero(signature.ReflectType())
	if n1 == n2 {
		// an existing method was overwritten.
		// it may be cached in some other type's method cache.
		t.universe.InvalidateMethodCache()
	}
	return index
}

// GetMethods returns the pointer to the method values.
// It panics if the type is unnamed
func (t *xtype) GetMethods() *[]r.Value {
	if !etoken.GENERICS.V2_CTI() && !t.Named() {
		xerrorf(t, "GetMethods on unnamed type %v", t)
	}
	resizemethodvalues(t, t.NumAllMethod())
	return &t.methodvalue
}
