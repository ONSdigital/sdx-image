// Package test provides helper functions for use within tests.
package test

import (
	"os"
	"path"
	"runtime"
)

// SetCwdToRoot sets the current working directory to the project root.
func SetCwdToRoot() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
