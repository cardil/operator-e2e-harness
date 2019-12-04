package test

import (
	"github.com/pkg/errors"
	"os"
	"os/user"
	"path"
	"strings"
	"testing"
)

var Kubeconfigs = calculate()
var PerformCleanup = func(t *testing.T) bool {
	return true
}

func calculate() []string {
	value, ok := os.LookupEnv("E2E_KUBECONFIGS")
	if !ok {
		value = defaultKubeconfig()
	}
	paths := strings.Split(value, ",")
	validatePaths(paths)
	return paths
}

func validatePaths(paths []string) {
	for _, p := range paths {
		_, err := os.Stat(p)
		if err != nil {
			panic(errors.WithMessagef(
				err, "All kubeconfig paths should exists and be readable. %v is not", p))
		}
	}
}

func defaultKubeconfig() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return path.Join(usr.HomeDir, ".kube", "config")
}
