package main

import (
	"flag"
	"log"

	"github.com/canonical/app-charm-config/internal/paascharm"
)

const (
	defaultCharmcraftLocation = "charmcraft.yaml"
	defaultPackageName        = "appconfig"
	defaultOutputFile         = "appconfig.go"
)

func main() {
	charmcraftFile := flag.String("c", defaultCharmcraftLocation, "charmcraft.yaml file location.")
	packageName := flag.String("p", defaultPackageName, "name of the generated package.")
	outputFile := flag.String("o", defaultOutputFile, "output file. Overwrites the previous file if it exists")
	flag.Parse()

	err := paascharm.CreateGoStructs(*charmcraftFile, *packageName, *outputFile)
	if err != nil {
		log.Fatal(err)
	}
}
