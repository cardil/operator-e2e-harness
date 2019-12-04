package catalogsource

import "log"

func Generate() string {
	log.Print("Generating catalog source...")
	return execute(execpath())
}
