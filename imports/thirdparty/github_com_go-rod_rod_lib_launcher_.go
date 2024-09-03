// this file was generated by gomacro command: import _i "github.com/go-rod/rod/lib/launcher"
// DO NOT EDIT! Any change will be lost when the file is re-generated

package launcher

import (
	r "reflect"
	launcher_ "github.com/go-rod/rod/lib/launcher"
)

// reflection: allow interpreted code to import "github.com/go-rod/rod/lib/launcher"
func init() {
	Packages["github.com/go-rod/rod/lib/launcher"] = Package{
		Name: "launcher",
		Binds: map[string]r.Value{
			"DefaultBrowserDir":	r.ValueOf(&launcher_.DefaultBrowserDir).Elem(),
			"DefaultUserDataDirPrefix":	r.ValueOf(&launcher_.DefaultUserDataDirPrefix).Elem(),
			"ErrAlreadyLaunched":	r.ValueOf(&launcher_.ErrAlreadyLaunched).Elem(),
			"HeaderName":	r.ValueOf(launcher_.HeaderName),
			"HostGoogle":	r.ValueOf(launcher_.HostGoogle),
			"HostNPM":	r.ValueOf(launcher_.HostNPM),
			"HostPlaywright":	r.ValueOf(launcher_.HostPlaywright),
			"LookPath":	r.ValueOf(launcher_.LookPath),
			"MustNewManaged":	r.ValueOf(launcher_.MustNewManaged),
			"MustResolveURL":	r.ValueOf(launcher_.MustResolveURL),
			"New":	r.ValueOf(launcher_.New),
			"NewAppMode":	r.ValueOf(launcher_.NewAppMode),
			"NewBrowser":	r.ValueOf(launcher_.NewBrowser),
			"NewManaged":	r.ValueOf(launcher_.NewManaged),
			"NewManager":	r.ValueOf(launcher_.NewManager),
			"NewURLParser":	r.ValueOf(launcher_.NewURLParser),
			"NewUserMode":	r.ValueOf(launcher_.NewUserMode),
			"Open":	r.ValueOf(launcher_.Open),
			"ResolveURL":	r.ValueOf(launcher_.ResolveURL),
			"RevisionDefault":	r.ValueOf(launcher_.RevisionDefault),
			"RevisionPlaywright":	r.ValueOf(launcher_.RevisionPlaywright),
		}, Types: map[string]r.Type{
			"Browser":	r.TypeOf((*launcher_.Browser)(nil)).Elem(),
			"Host":	r.TypeOf((*launcher_.Host)(nil)).Elem(),
			"Launcher":	r.TypeOf((*launcher_.Launcher)(nil)).Elem(),
			"Manager":	r.TypeOf((*launcher_.Manager)(nil)).Elem(),
			"URLParser":	r.TypeOf((*launcher_.URLParser)(nil)).Elem(),
		}, Untypeds: map[string]string{
			"HeaderName":	"string:Rod-Launcher",
			"RevisionDefault":	"int:1321438",
			"RevisionPlaywright":	"int:1124",
		}, 
	}
}
