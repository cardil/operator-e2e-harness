package catalogsource

import (
	"bytes"
	"github.com/pkg/errors"
	"log"
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

func execute(command string) {
	out, errout := capture(command)
	printOut(out)
	printErr(errout)
}

func capture(command string) (string, string) {
	log.Printf("Running command: %s", command)
	err, out, errout := shell(command)
	if err != nil {
		printOut(out)
		printErr(errout)
		panic(errors.WithMessagef(err, "Error while running command: %s", command))
	}
	return out, errout
}

func printOut(out string) {
	if out != "" {
		log.Printf("--- stdout ---")
		log.Print(out)
	}
}

func printErr(errout string) {
	if errout != "" {
		log.Print("--- stderr ---")
		log.Print(errout)
	}
}
