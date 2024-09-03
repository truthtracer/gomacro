package genimport

import (
	"fmt"
	"go/build"
	"os"
	"strings"
)

const sep = string(os.PathSeparator)

// GoGenerateMain allows gomacro to be run under
// go generate. It is used to write new x_package.go
// import bindings for a package. Thus `go generate`
// will automatically update your bindings.
//
// To use, add a comment to one go file in your package:
//
// `//go:generate gomacro -g .`
//
// to import the current dir; or one like
//
// `//go:generate gomacro -g github.com/truthtracer/gomacro/classic`
//
// to specify the exact import path. The second, specific
// form, may be necessary if we cannot detect the GOPATH
// environment variable.
func GoGenerateMain(arg []string, imp *Importer) error {
	var path string
	narg := len(arg)
	switch {
	case narg == 0 || (narg > 0 && arg[0] == "."):
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("gomacro -g: error getting current dir: %v", err)
		}
		gopath := build.Default.GOPATH
		prefix := gopath + sep + "src" + sep
		if strings.HasPrefix(cwd, prefix) {
			path = cwd[len(prefix):]
		} else {
			// guess it is after the first `src` in cwd,
			// since traditionally all packages are
			// after $GOPATH/src/
			splt := strings.SplitN(cwd, sep+"src"+sep, 2)
			if len(splt) <= 1 {
				return fmt.Errorf("gomacro -g: unable to detect current package, please specify it")
			}
			path = splt[1]
		}
	default:
		path = arg[0]
		if isLocalFilesystemPath(path) {
			abspath, err := MakeAbsolutePathOrError(path)
			if err != nil {
				return err
			}
			path = abspath.String()
		}
	}
	_, err := imp.ImportPackagesOrError(
		map[string]PackageName{path: PackageName("_i")},
		true /*enableModule*/)
	return err
}
