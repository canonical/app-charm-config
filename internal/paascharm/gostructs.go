package paascharm

import (
	"fmt"
	"log"
	"sort"
	"strings"
	"unicode"
)

const CommonPrefix = "APP_"

// Map from integration key in charmcraft.yaml to Go name
var integrationsToGoName = map[string]string{
	"mongodb":    "MongoDB",
	"mysql":      "MySQL",
	"postgresql": "PostgreSQL",
	"redis":      "Redis",
	"s3":         "S3",
	"saml":       "SAML",
}

// Map from database integration key in charmcraft.yaml to its prefix in env vars
var databaseIntegrationNameToPrefix = map[string]string{
	"mongodb":    CommonPrefix + "MONGODB_",
	"mysql":      CommonPrefix + "MYSQL_",
	"postgresql": CommonPrefix + "POSTGRESQL_",
	"redis":      CommonPrefix + "REDIS_",
}

// Charmcraft config options types to Go Types
var charmcraftToGoTypes = map[string]string{
	"bool":    "bool",
	"boolean": "bool",
	"float":   "float64",
	"int":     "int",
	"secret":  "string",
	"string":  "string",
}

type GoStructsData struct {
	PackageName             string
	CommonPrefix            string
	Options                 []Option
	HasDatabaseIntegrations bool
	Integrations            map[string]Integration
}

type Option struct {
	GoVarName  string
	GoVarType  string
	EnvVarName string
}

type Integration struct {
	Name           string
	GoName         string
	Interface      string
	Optional       bool
	IsDatabase     bool
	DatabasePrefix string
}

// Create a GoStructsData with all the information needed to generate the Go file from the charmcraft.yaml parsed file
func NewGoStructsData(packageName string, charmcraft CharmcraftYAMLConfig) (GoStructsData, error) {
	goStructs := GoStructsData{
		PackageName:  packageName,
		CommonPrefix: CommonPrefix,
	}

	for key, value := range charmcraft.Config.Options {
		varType, err := buildGoVarType(value)
		if err != nil {
			return GoStructsData{}, err
		}
		option := Option{
			GoVarName:  buildGoVarName(key),
			GoVarType:  varType,
			EnvVarName: buildEnvVarName(key),
		}
		goStructs.Options = append(goStructs.Options, option)
	}

	goStructs.Integrations = make(map[string]Integration)
	for key, value := range charmcraft.Requires {
		integration := Integration{
			Name:      key,
			Interface: value.Interface,
			Optional:  value.Optional,
		}

		goName, ok := integrationsToGoName[key]
		if !ok {
			log.Printf("Skipping unknown integration %s\n", key)
			continue
		}
		integration.GoName = goName
		if databasePrefix, okDatabase := databaseIntegrationNameToPrefix[key]; okDatabase {
			goStructs.HasDatabaseIntegrations = true
			integration.IsDatabase = true
			integration.DatabasePrefix = databasePrefix
		}
		goStructs.Integrations[key] = integration
	}

	return normalise(goStructs), nil
}

// config option name to Go variable name.
// capitalises first letter and every letter after a '-'
// and also removes '-'. user-config will become UserConfig.
func buildGoVarName(configName string) (result string) {
	parts := strings.Split(configName, "-")
	for _, part := range parts {
		partRunes := []rune(part)
		if len(partRunes) > 0 {
			partRunes[0] = unicode.ToUpper(partRunes[0])
			result += string(partRunes)
		}
	}
	return result
}

func buildGoVarType(configOption CharmcraftConfigOption) (string, error) {
	goType, ok := charmcraftToGoTypes[configOption.Type]
	if !ok {
		return "", fmt.Errorf("unknown type for config option of type: %s", configOption.Type)
	}
	result := goType

	// If there is no default value for a config option, a pointer can help differentiate between
	// the empty value and no value specified in the config option.
	if configOption.Default == nil {
		result = "*" + result
	}
	return result, nil
}

func buildEnvVarName(configName string) string {
	result := CommonPrefix + configName
	result = strings.ReplaceAll(result, "-", "_")
	result = strings.ToUpper(result)
	return result
}

// Normalise the GoStructsData, so the same input always outputs the same output.
// Specifically, order the Options slice by GoVarName, as the order was given by
// iterating over a map.
func normalise(goStructsData GoStructsData) GoStructsData {
	result := goStructsData
	orderedOptions := make([]Option, len(result.Options))
	copy(orderedOptions, result.Options)
	sort.Slice(result.Options, func(i, j int) bool {
		return orderedOptions[i].GoVarName < orderedOptions[j].GoVarName
	})
	return result
}
