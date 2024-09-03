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
 * convert.go
 *
 *  Created on Apr 30, 2017
 *      Author Massimiliano Ghilardi
 */

package fast

import (
	"go/ast"
	r "reflect"

	"github.com/truthtracer/gomacro/base/reflect"

	xr "github.com/truthtracer/gomacro/xreflect"
)

// Convert compiles a type conversion expression
func (c *Comp) Convert(node ast.Expr, t xr.Type) *Expr {
	e := c.expr1(node, nil)
	return c.convert(e, t, node)
}

// Convert compiles a type conversion expression
func (c *Comp) convert(e *Expr, t xr.Type, nodeOpt ast.Expr) *Expr {
	if e.Untyped() {
		e.ConstTo(t)
	}

	if e.Type != nil && e.Type.IdenticalTo(t) {
		return e
	} else if e.Type != nil && e.Type.ReflectType() == t.ReflectType() {
		if e.Const() {
			return c.exprValue(t, e.Value)
		} else {
			return exprFun(t, e.Fun)
		}
	} else if e.Type == nil && reflect.IsNillableKind(t.Kind()) {
		e.Type = t
		e.Value = xr.Zero(t).Interface()
	} else if e.Type != nil && e.Type.ConvertibleTo(t) {
	} else {
		c.Errorf("cannot convert %v to %v: %v", e.Type, t, nodeOpt)
		return nil
	}
	rtype := t.ReflectType()
	if e.Const() {
		val := convert(xr.ValueOf(e.Value), rtype).Interface()
		return c.exprValue(t, val)
	}
	fun := e.AsX1()
	var ret I
	switch t.Kind() {
	case xr.Bool:
		ret = func(env *Env) bool {
			val := convert(fun(env), rtype)
			return val.Bool()
		}
	case xr.Int:
		ret = func(env *Env) int {
			val := convert(fun(env), rtype)
			return int(val.Int())
		}
	case xr.Int8:
		ret = func(env *Env) int8 {
			val := convert(fun(env), rtype)
			return int8(val.Int())
		}
	case xr.Int16:
		ret = func(env *Env) int16 {
			val := convert(fun(env), rtype)
			return int16(val.Int())
		}
	case xr.Int32:
		ret = func(env *Env) int32 {
			val := convert(fun(env), rtype)
			return int32(val.Int())
		}
	case xr.Int64:
		ret = func(env *Env) int64 {
			val := convert(fun(env), rtype)
			return val.Int()
		}
	case xr.Uint:
		ret = func(env *Env) uint {
			val := convert(fun(env), rtype)
			return uint(val.Uint())
		}
	case xr.Uint8:
		ret = func(env *Env) uint8 {
			val := convert(fun(env), rtype)
			return uint8(val.Uint())
		}
	case xr.Uint16:
		ret = func(env *Env) uint16 {
			val := convert(fun(env), rtype)
			return uint16(val.Uint())
		}
	case xr.Uint32:
		ret = func(env *Env) uint32 {
			val := convert(fun(env), rtype)
			return uint32(val.Uint())
		}
	case xr.Uint64:
		ret = func(env *Env) uint64 {
			val := convert(fun(env), rtype)
			return val.Uint()
		}
	case xr.Uintptr:
		ret = func(env *Env) uintptr {
			val := convert(fun(env), rtype)
			return uintptr(val.Uint())
		}
	case xr.Float32:
		ret = func(env *Env) float32 {
			val := convert(fun(env), rtype)
			return float32(val.Float())
		}
	case xr.Float64:
		ret = func(env *Env) float64 {
			val := convert(fun(env), rtype)
			return val.Float()
		}
	case xr.Complex64:
		ret = func(env *Env) complex64 {
			val := convert(fun(env), rtype)
			return complex64(val.Complex())
		}
	case xr.Complex128:
		ret = func(env *Env) complex128 {
			val := convert(fun(env), rtype)
			return val.Complex()
		}
	case xr.String:
		ret = func(env *Env) string {
			val := convert(fun(env), rtype)
			return val.String()
		}
	default:
		if conv := c.Converter(e.Type, t); conv != nil {
			ret = func(env *Env) xr.Value {
				return conv(fun(env))
			}
		} else {
			ret = func(env *Env) xr.Value {
				return fun(env)
			}
		}
	}
	eret := exprFun(t, ret)
	if e.Const() {
		eret.EvalConst(COptKeepUntyped)
	}
	return eret
}

// Converter returns a function that converts reflect.Value from tin to tout
// also supports conversion from interpreted types to interfaces
func (c *Comp) Converter(tin, tout xr.Type) func(xr.Value) xr.Value {
	if !tin.ConvertibleTo(tout) {
		c.Errorf("cannot convert from <%v> to <%v>", tin, tout)
	}
	rtin := tin.ReflectType()
	rtout := tout.ReflectType()
	switch {
	case rtin == rtout:
		return nil
	case rtin.ConvertibleTo(rtout):
		// most conversions, including from compiled type to compiled interface
		if rtin.Kind() != r.Interface {
			return func(obj xr.Value) xr.Value {
				return obj.Convert(rtout)
			}
		}
		// extract objects wrapped in proxies (if any)
		g := c.CompGlobals
		return func(obj xr.Value) xr.Value {
			obj, _ = g.extractFromProxy(obj)
			if obj.IsValid() {
				return obj.Convert(rtout)
			} else {
				return xr.ZeroR(rtout)
			}
		}
	case xr.IsEmulatedInterface(tout):
		// conversion from type to emulated interface
		return c.converterToEmulatedInterface(tin, tout)
	case rtin == rtypeOfForward:
		// conversion from forward-declared type
		return c.converterFromForward(tin, tout)
	case rtout.Kind() == r.Interface:
		// conversion from interpreted type to compiled interface.
		// must use a proxy that pre-implement compiled interfaces.
		return c.converterToProxy(tin, tout)
	default:
		c.Errorf("unimplemented conversion from <%v> to <%v> with reflect.Type <%v> to <%v>",
			tin, tout, rtin, rtout)
		return nil
	}
}

// conversion from forward-declared type
func (c *Comp) converterFromForward(tin, tout xr.Type) func(xr.Value) xr.Value {
	rtout := tout.ReflectType()
	return func(val xr.Value) xr.Value {
		val = val.Elem()
		if val.Type() != rtout {
			val = val.Convert(rtout)
		}
		return val
	}
}

// conversion between compatible types.
// also implements conversion from xr.Forward.
func convert(v xr.Value, rtout r.Type) xr.Value {
	if v.Kind() == r.Interface {
		v = v.Elem()
	}
	return v.Convert(rtout)
}
