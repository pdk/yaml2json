// Read stdin expecting a yaml file, and write data as JSON on stdout.

package main

import (
	"log"
	"os"

	"github.com/pdk/yaml2json"
)

func main() {

	t, err := yaml2json.ReadYAML(os.Stdin)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = yaml2json.PrintJSON(os.Stdout, t)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
