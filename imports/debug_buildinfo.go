// this file was generated by gomacro command: import _b "debug/buildinfo"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package imports

import (
	. "reflect"
	buildinfo "debug/buildinfo"
)

// reflection: allow interpreted code to import "debug/buildinfo"
func init() {
	Packages["debug/buildinfo"] = Package{
		Name: "buildinfo",
		Binds: map[string]Value{
			"Read":	ValueOf(buildinfo.Read),
			"ReadFile":	ValueOf(buildinfo.ReadFile),
		}, Types: map[string]Type{
			"BuildInfo":	TypeOf((*buildinfo.BuildInfo)(nil)).Elem(),
		}, 
	}
}
