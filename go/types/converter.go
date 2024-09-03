// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file converts objects from go/types to github.com/truthtracer/go/types

package types

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/types/typeutil"
)

type Converter struct {
	pkg          map[string]*Package
	cache        typeutil.Map
	toaddmethods map[*Named]*types.Named
	tocomplete   []*Interface
}

type funcOption bool

const (
	funcIgnoreRecv funcOption = false
	funcSetRecv    funcOption = true

	debugConverter = false // set to true to enable debug messages
)

// should be called with argument Universe
// to initialize basic types, constants true/false/iota and 'error'
func (c *Converter) Init(universe *Scope) {
	// create builtin package i.e. universe
	p := NewPackage("", "")
	if universe != nil {
		// fill package with contents of universe scope
		scope := p.Scope()
		for _, name := range universe.Names() {
			scope.Insert(universe.Lookup(name))
		}
	}
	c.pkg = map[string]*Package{
		"": p,
	}
}

// convert *go/types.Package -> *github.com/truthtracer/gomacro/go/types.Package
func (c *Converter) Package(g *types.Package) *Package {
	if g == nil {
		return nil
	}
	c.cache = typeutil.Map{}
	p := c.mkpackage(g)
	scope := g.Scope()
	for _, name := range scope.Names() {
		obj := c.object(scope.Lookup(name))
		if obj != nil {
			p.scope.Insert(obj)
		}
	}
	// complete all interfaces
	for _, t := range c.tocomplete {
		t.Complete()
	}
	c.tocomplete = c.tocomplete[0:0:cap(c.tocomplete)]

	// add methods to named types
	for t, g := range c.toaddmethods {
		c.addmethods(t, g)
		delete(c.toaddmethods, t)
	}

	return p
}

// convert go/types.Object -> github.com/truthtracer/gomacro/go/types.Object
func (c *Converter) object(g types.Object) (ret Object) {
	defer func() {
		if ret == nil {
			if e := recover(); e != nil {
				fmt.Printf("// warning: skipping import of %v:\t%v\n", g, e)
			}
		}
	}()
	switch g := g.(type) {
	case *types.Const:
		ret = c.constant(g)
	case *types.Func:
		ret = c.function(g)
	case *types.TypeName:
		ret = c.typename(g)
	case *types.Var:
		ret = c.variable(g)
	default:
	}
	return ret
}

// convert *go/types.Const -> *github.com/truthtracer/gomacro/go/types.Const
func (c *Converter) constant(g *types.Const) *Const {
	return NewConst(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()), g.Val())
}

// convert *go/types.Func -> *github.com/truthtracer/gomacro/go/types.Func
func (c *Converter) function(g *types.Func) *Func {
	return NewFunc(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()).(*Signature))
}

// convert *go/types.TypeName -> *github.com/truthtracer/gomacro/go/types.TypeName
func (c *Converter) typename(g *types.TypeName) *TypeName {
	ret, _ := c.mktypename(g)
	if ret.typ == nil {
		ret.typ = c.typ(g.Type())
	}
	return ret
}

// convert *go/types.Var -> *github.com/truthtracer/gomacro/go/types.Var
func (c *Converter) variable(g *types.Var) *Var {
	return NewVar(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()))
}

func (c *Converter) typ(g types.Type) Type {
	if g == nil {
		return nil
	}
	t, _ := c.cache.At(g).(Type)
	if t != nil {
		return t
	}
	switch g := g.(type) {
	case *types.Array:
		elem := c.typ(g.Elem())
		t = NewArray(elem, g.Len())
	case *types.Basic:
		return Typ[BasicKind(g.Kind())]
	case *types.Chan:
		elem := c.typ(g.Elem())
		t = NewChan(ChanDir(g.Dir()), elem)
	case *types.Interface:
		t = c.mkinterface(g)
	case *types.Map:
		t = c.mkmap(g)
	case *types.Named:
		t = c.mknamed(g)
	case *types.Pointer:
		elem := c.typ(g.Elem())
		t = NewPointer(elem)
	case *types.Signature:
		t = c.mksignature(g, funcSetRecv)
	case *types.Slice:
		elem := c.typ(g.Elem())
		t = NewSlice(elem)
	case *types.Struct:
		t = c.mkstruct(g)
	case *types.TypeParam:
		panic(fmt.Errorf("importing generic functions or types is not supported yet"))
	default:
		panic(fmt.Errorf("Converter.Type(): unsupported types.Type: %T", g))
	}
	c.cache.Set(g, t)
	if debugConverter {
		fmt.Print("scanned type ", t, " has ", t.NumMethods(), " methods")
		if u := t.Underlying(); u != nil {
			fmt.Print(", and its underlying type ", u, " has ", u.NumMethods(), " methods")
		}
		fmt.Println()
	}
	return t
}

func (c *Converter) mkinterface(g *types.Interface) *Interface {
	n := g.NumExplicitMethods()
	fs := make([]*Func, n)
	for i := 0; i < n; i++ {
		fs[i] = c.mkfunc(g.ExplicitMethod(i), funcIgnoreRecv)
		if debugConverter {
			fmt.Println("added interface method", fs[i].Name())
		}
	}
	n = g.NumEmbeddeds()
	es := make([]Type, n)
	for i := 0; i < n; i++ {
		es[i] = c.typ(g.EmbeddedType(i))
		if debugConverter {
			fmt.Println("added embedded interface", es[i])
		}
	}
	t := NewInterfaceType(fs, es)
	c.tocomplete = append(c.tocomplete, t)
	return t
}

func (c *Converter) mkmap(g *types.Map) *Map {
	key := c.typ(g.Key())
	elem := c.typ(g.Elem())
	return NewMap(key, elem)
}

func (c *Converter) mknamed(g *types.Named) *Named {
	typename, found := c.mktypename(g.Obj())
	if found && typename.Type() != nil {
		return typename.Type().(*Named)
	}
	t := NewNamed(typename, nil, nil)
	// cache t early, in case it's part of a cycle in a recursive type
	c.cache.Set(g, t)
	if debugConverter {
		fmt.Println("scanning underlying type of", typename.Name())
	}
	u := c.typ(g.Underlying())
	t.SetUnderlying(u)
	if g.NumMethods() != 0 {
		if c.toaddmethods == nil {
			c.toaddmethods = make(map[*Named]*types.Named)
		}
		c.toaddmethods[t] = g
	}
	return t
}

func (c *Converter) mksignature(g *types.Signature, opt funcOption) *Signature {
	var recv *Var
	if opt == funcSetRecv {
		recv = c.mkparam(g.Recv())
	}
	return NewSignature(
		recv,
		c.mkparams(g.Params()),
		c.mkparams(g.Results()),
		g.Variadic(),
	)
}

func (c *Converter) mkstruct(g *types.Struct) *Struct {
	n := g.NumFields()
	fields := make([]*Var, n)
	tags := make([]string, n)
	for i := 0; i < n; i++ {
		fields[i] = c.mkfield(g.Field(i))
		tags[i] = g.Tag(i)
	}
	return NewStruct(fields, tags)
}

func (c *Converter) mkpackage(g *types.Package) *Package {
	if g == nil {
		return nil
	}
	path := g.Path()
	if p := c.pkg[path]; p != nil {
		return p
	}
	p := NewPackage(path, g.Name())
	c.pkg[path] = p
	return p
}

func (c *Converter) universe() *Package {
	return c.pkg[""]
}

func (c *Converter) mktypename(g *types.TypeName) (*TypeName, bool) {
	pkg := c.mkpackage(g.Pkg())
	if pkg == nil {
		pkg = c.universe()
	}
	scope := pkg.Scope()
	obj := scope.Lookup(g.Name())
	// to preserve type identity, reuse existing typename if found
	if typename, ok := obj.(*TypeName); ok {
		return typename, true
	}
	typename := NewTypeName(g.Pos(), pkg, g.Name(), nil)
	pkg.Scope().Insert(typename)
	return typename, false
}

func (c *Converter) mkfield(g *types.Var) *Var {
	// g.Embedded() is a newer alias for g.Anonymous(),
	// but go 1.9 does not have it
	return NewField(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()), g.Anonymous())
}

func (c *Converter) mkparam(g *types.Var) *Var {
	if g == nil {
		return nil
	}
	return NewParam(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()))
}

func (c *Converter) mkparams(g *types.Tuple) *Tuple {
	if g == nil {
		return nil
	}
	n := g.Len()
	v := make([]*Var, n)
	for i := 0; i < n; i++ {
		v[i] = c.mkparam(g.At(i))
	}
	return NewTuple(v...)
}

func (c *Converter) mkvar(g *types.Var) *Var {
	if g == nil {
		return nil
	}
	return NewVar(g.Pos(), c.mkpackage(g.Pkg()), g.Name(), c.typ(g.Type()))
}

func (c *Converter) mkfunc(m *types.Func, opt funcOption) *Func {
	sig := c.mksignature(m.Type().(*types.Signature), opt)
	return NewFunc(m.Pos(), c.mkpackage(m.Pkg()), m.Name(), sig)
}

func (c *Converter) addmethods(t *Named, g *types.Named) {
	n := g.NumMethods()
	for i := 0; i < n; i++ {
		m := c.mkfunc(g.Method(i), funcSetRecv)
		t.AddMethod(m)
	}
}
