package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/canonical/app-charm-config/internal/paascharm"
)

const (
	defaultCharmcraftLocation = "charmcraft.yaml"
	defaultPackageName        = "appconfig"
	defaultOutputFile         = "appconfig.go"
)

var (
	genCmd         = flag.NewFlagSet("gen", flag.ExitOnError)
	charmcraftFile = genCmd.String("c", defaultCharmcraftLocation, "charmcraft.yaml file location.")
	packageName    = genCmd.String("p", defaultPackageName, "name of the generated package.")
	outputFile     = genCmd.String("o", defaultOutputFile, "output file. Overwrites the previous file if it exists")
)

func main() {
	genCmd.Usage = usage
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		flag.Usage()
		os.Exit(2)
	}

	subCmd, subCmdArgs := args[0], args[1:]
	switch subCmd {
	case genCmd.Name():
		genCmd.Parse(subCmdArgs)
		err := paascharm.CreateGoStructs(*charmcraftFile, *packageName, *outputFile)
		if err != nil {
			log.Fatal(err)
		}
	default:
		flag.Usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s %s [command flags]\n", os.Args[0], genCmd.Name())
	fmt.Fprintf(flag.CommandLine.Output(), "Where [command flags] for %s subcommand can be:\n", genCmd.Name())
	genCmd.PrintDefaults()
}
