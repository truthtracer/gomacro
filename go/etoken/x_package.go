// this file was generated by gomacro command: import _i "github.com/truthtracer/gomacro/go/etoken"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package etoken

import (
	r "reflect"

	"github.com/truthtracer/gomacro/imports"
)

// reflection: allow interpreted code to import "github.com/truthtracer/gomacro/go/etoken"
func init() {
	imports.Packages["github.com/truthtracer/gomacro/go/etoken"] = imports.Package{
		Binds: map[string]r.Value{
			"FUNCTION":       r.ValueOf(FUNCTION),
			"IsKeyword":      r.ValueOf(IsKeyword),
			"IsLiteral":      r.ValueOf(IsLiteral),
			"IsMacroKeyword": r.ValueOf(IsMacroKeyword),
			"IsOperator":     r.ValueOf(IsOperator),
			"LAMBDA":         r.ValueOf(LAMBDA),
			"Lookup":         r.ValueOf(Lookup),
			"LookupSpecial":  r.ValueOf(LookupSpecial),
			"MACRO":          r.ValueOf(MACRO),
			"NewFileSet":     r.ValueOf(NewFileSet),
			"QUASIQUOTE":     r.ValueOf(QUASIQUOTE),
			"QUOTE":          r.ValueOf(QUOTE),
			"String":         r.ValueOf(String),
			"TYPECASE":       r.ValueOf(TYPECASE),
			"UNQUOTE":        r.ValueOf(UNQUOTE),
			"UNQUOTE_SPLICE": r.ValueOf(UNQUOTE_SPLICE),
		},
		Types: map[string]r.Type{
			"File":    r.TypeOf((*File)(nil)).Elem(),
			"FileSet": r.TypeOf((*FileSet)(nil)).Elem(),
			"Token":   r.TypeOf((*Token)(nil)).Elem(),
		},
		Proxies: map[string]r.Type{}}
}
