package catalogsource

import "log"

func Generate() string {
	log.Print("Generating catalog source...")
	out, _ := capture(execpath())
	return out
}
