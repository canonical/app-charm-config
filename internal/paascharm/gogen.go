package paascharm

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/format"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

// CreateGoStructs reads the charmcraft.yaml file from the charmcraftFile location
// and outputs the Go file with the proper structs in the outputFile location,
// using as package name packageName. It will override the outputFile file if
// it exists and create all the parent directories if missing.
func CreateGoStructs(charmcraftFileName, packageName, outputFile string) error {
	yamlFile, err := os.Open(charmcraftFileName)
	if err != nil {
		return fmt.Errorf("cannot read charmcraft.yaml file: %v", err)
	}
	defer yamlFile.Close()

	charmcraftConfig, err := ParseCharmcraftYAML(yamlFile)
	if err != nil {
		return fmt.Errorf("cannot parse charmcraft.yaml file: %v", err)
	}

	goStructInfo, err := NewGoStructsData(packageName, charmcraftConfig)
	if err != nil {
		return fmt.Errorf("cannot create go structs data: %v", err)
	}

	goStructs, err := GenerateGoStructs(goStructInfo)
	if err != nil {
		return fmt.Errorf("cannot generate go structs info: %v", err)
	}

	err = os.MkdirAll(filepath.Dir(outputFile), 0755)
	if err != nil {
		return fmt.Errorf("cannot create parent directories for output file %s: %v", outputFile, err)
	}

	if _, err := os.Stat(outputFile); err == nil {
		log.Printf("file %s already exists and will be overridden\n", outputFile)
	}

	err = os.WriteFile(outputFile, goStructs, 0644)
	if err != nil {
		return fmt.Errorf("cannot write output file %s: %v", outputFile, err)
	}
	log.Printf("configuration written to file: %s\n", outputFile)

	return nil
}

//go:embed go.tmpl
var goTemplateSource string

var goTemplate = template.Must(template.New("").Parse(goTemplateSource))

// Generate a []byte with the Go file containing the Go structs for a GoStructsData struct.
// The output code is formatted following gofmt style.
func GenerateGoStructs(goStructsData GoStructsData) ([]byte, error) {
	var buf bytes.Buffer
	err := goTemplate.Execute(&buf, goStructsData)
	if err != nil {
		return nil, fmt.Errorf("cannot execute go template: %v", err)
	}

	goStructs, err := format.Source(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("cannot format go code: %v", err)
	}

	return goStructs, nil
}
