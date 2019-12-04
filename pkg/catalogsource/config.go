package catalogsource

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
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
	rootdir := path.Dir(path.Dir(wd))
	catalogPath := path.Join(rootdir, "hack", "catalog.sh")
	log.Printf("Default catalog source generator script path: %s", catalogPath)
	return catalogPath
}

func execpath() string {
	_, err := os.Stat(Executable)
	if err != nil {
		panic(errors.WithMessage(err, "CatalogSource generator executable file should be present"))
	}
	return Executable
}
