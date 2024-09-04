package paascharm_test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/canonical/app-charm-config/internal/paascharm"
)

func TestCreateGoStructsWithMinimalCharmcraftYAML(t *testing.T) {
	tmpdir := t.TempDir()
	charmcraftFileName := path.Join(tmpdir, "charmcraft.yaml")
	err := os.WriteFile(charmcraftFileName, []byte("name: go-app\n"), 0644)
	if err != nil {
		t.Errorf("Error writing charmcraft.yaml file: %v\n", err)
	}
	packageName := "myconfigcharm"
	outputFile := path.Join(tmpdir, packageName, "config.go")

	err = paascharm.CreateGoStructs(charmcraftFileName, packageName, outputFile)
	if err != nil {
		t.Errorf("Error creating go Structs: %v\n", err)
	}

	// Check that the outputfile exists and containes the correct package name
	data, err := os.ReadFile(outputFile)
	if err != nil {
		t.Error("Error reading output file")
	}
	strData := string(data)
	if !strings.Contains(strData, "package myconfigcharm") {
		t.Error("Output file does not contain package name")
	}
}
