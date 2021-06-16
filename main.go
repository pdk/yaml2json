// Read stdin expecting a yaml file, and rewrite as json to stdout.
package main

import (
	"log"
	"os"

	"github.com/pdk/yaml2json/rwc"
)

func main() {

	t, err := rwc.ReadYAML(os.Stdin)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = rwc.PrintJSON(os.Stdout, t)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
