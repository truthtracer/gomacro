// this file was generated by gomacro command: import _b "go/build/constraint"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	constraint "go/build/constraint"
)

// reflection: allow interpreted code to import "go/build/constraint"
func init() {
	Packages["go/build/constraint"] = Package{
		Name: "constraint",
		Binds: map[string]Value{
			"IsGoBuild":	ValueOf(constraint.IsGoBuild),
			"IsPlusBuild":	ValueOf(constraint.IsPlusBuild),
			"Parse":	ValueOf(constraint.Parse),
			"PlusBuildLines":	ValueOf(constraint.PlusBuildLines),
		}, Types: map[string]Type{
			"AndExpr":	TypeOf((*constraint.AndExpr)(nil)).Elem(),
			"Expr":	TypeOf((*constraint.Expr)(nil)).Elem(),
			"NotExpr":	TypeOf((*constraint.NotExpr)(nil)).Elem(),
			"OrExpr":	TypeOf((*constraint.OrExpr)(nil)).Elem(),
			"SyntaxError":	TypeOf((*constraint.SyntaxError)(nil)).Elem(),
			"TagExpr":	TypeOf((*constraint.TagExpr)(nil)).Elem(),
		}, 
	}
}
