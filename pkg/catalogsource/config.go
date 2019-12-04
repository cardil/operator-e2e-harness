package catalogsource

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
)

var Executable = calculateDefaultPath()
var rundir = (func() string {
	dir, err := ioutil.TempDir("", "e2e-harness")
	if err != nil {
		panic(errors.WithMessage(err, "Can't get temporary directory"))
	}
	return dir
})()

func calculateDefaultPath() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(errors.WithMessage(err, "Can't get working directory"))
	}
	return path.Join(wd, "hack", "catalog.sh")
}

func execpath() string {
	_, err := os.Stat(Executable)
	if err != nil {
		panic(errors.WithMessage(err, "CatalogSource generator executable file should be present"))
	}
	return Executable
}
