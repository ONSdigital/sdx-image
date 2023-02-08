package test

import (
	"os"
	"path"
	"runtime"
)

func SetCwdToRoot() {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "../..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}
}
