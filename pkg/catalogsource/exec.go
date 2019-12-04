package catalogsource

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"os"
	"os/exec"
)

func shell(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func execute(command string) string {
	fmt.Printf("Running command: %s", command)
	err, out, errout := shell(command)
	if errout != "" {
		_, _ = fmt.Fprint(os.Stderr, "--- stderr ---")
		_, _ = fmt.Fprint(os.Stderr, errout)
	}
	if err != nil {
		panic(errors.WithMessagef(err, "Error while running command: %s", command))
	}
	return out
}
