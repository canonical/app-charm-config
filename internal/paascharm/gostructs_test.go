package paascharm_test

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"

	"github.com/canonical/app-charm-config/internal/paascharm"
)

func TestConfigOptions(t *testing.T) {
	var tests = []struct {
		configOptionName       string
		charmcraftConfigOption paascharm.CharmcraftConfigOption
		expected               paascharm.Option
	}{{
		"string-config-option",
		paascharm.CharmcraftConfigOption{Type: "string"},
		paascharm.Option{"StringConfigOption", "*string", "APP_STRING_CONFIG_OPTION"},
	}, {
		"stringconfigoptionwithdefault",
		paascharm.CharmcraftConfigOption{Type: "string", Default: "def"},
		paascharm.Option{"Stringconfigoptionwithdefault", "string", "APP_STRINGCONFIGOPTIONWITHDEFAULT"},
	}, {
		"bool-config-option",
		paascharm.CharmcraftConfigOption{Type: "bool"},
		paascharm.Option{"BoolConfigOption", "*bool", "APP_BOOL_CONFIG_OPTION"},
	}, {
		"boolean-config-option-with-default",
		paascharm.CharmcraftConfigOption{Type: "boolean", Default: true},
		paascharm.Option{"BooleanConfigOptionWithDefault", "bool", "APP_BOOLEAN_CONFIG_OPTION_WITH_DEFAULT"},
	}, {
		"int-config-option",
		paascharm.CharmcraftConfigOption{Type: "int"},
		paascharm.Option{"IntConfigOption", "*int", "APP_INT_CONFIG_OPTION"},
	}, {
		"float-config-option",
		paascharm.CharmcraftConfigOption{Type: "float"},
		paascharm.Option{"FloatConfigOption", "*float64", "APP_FLOAT_CONFIG_OPTION"},
	}, {
		"secret-config-option",
		paascharm.CharmcraftConfigOption{Type: "secret"},
		paascharm.Option{"SecretConfigOption", "*string", "APP_SECRET_CONFIG_OPTION"},
	}}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.configOptionName)
		t.Run(testname, func(t *testing.T) {
			charmcraftYAMLConfig := paascharm.CharmcraftYAMLConfig{
				Config: paascharm.CharmcraftConfig{
					Options: map[string]paascharm.CharmcraftConfigOption{tt.configOptionName: tt.charmcraftConfigOption},
				},
			}
			goStructsData, err := paascharm.NewGoStructsData("packagename", charmcraftYAMLConfig)
			if err != nil {
				t.Fatalf("Error creating Go structs data %s", err)
			}

			if len(goStructsData.Options) != 1 {
				t.Fatalf("Wrong number of config optoins in Go structs data")
			}
			actual := goStructsData.Options[0]

			if diff := pretty.Diff(actual, tt.expected); len(diff) > 0 {
				t.Errorf("config optoin is not correctly generated. diff: %s\n", diff)
			}
		})
	}
}

func TestIntegrations(t *testing.T) {
	var tests = []struct {
		name           string
		charmcraftYAML paascharm.CharmcraftYAMLConfig
		expected       paascharm.GoStructsData
	}{{
		"mongodb mysql redis and postgresql",
		paascharm.CharmcraftYAMLConfig{
			Requires: map[string]paascharm.CharmcraftIntegration{
				"mongodb":    {Interface: "mongodb_client"},
				"mysql":      {Interface: "mysql_client"},
				"postgresql": {Interface: "postgresql_client"},
				"redis":      {Interface: "redis"},
			},
		},
		paascharm.GoStructsData{
			PackageName:             "pkg",
			CommonPrefix:            "APP_",
			HasDatabaseIntegrations: true,
			Integrations: map[string]paascharm.Integration{
				"mongodb":    {Name: "mongodb", GoName: "MongoDB", Interface: "mongodb_client", IsDatabase: true, DatabasePrefix: "APP_MONGODB_"},
				"mysql":      {Name: "mysql", GoName: "MySQL", Interface: "mysql_client", IsDatabase: true, DatabasePrefix: "APP_MYSQL_"},
				"postgresql": {Name: "postgresql", GoName: "PostgreSQL", Interface: "postgresql_client", IsDatabase: true, DatabasePrefix: "APP_POSTGRESQL_"},
				"redis":      {Name: "redis", GoName: "Redis", Interface: "redis", IsDatabase: true, DatabasePrefix: "APP_REDIS_"},
			},
		},
	}, {
		"optional s3 and saml, no database integration",
		paascharm.CharmcraftYAMLConfig{
			Requires: map[string]paascharm.CharmcraftIntegration{
				"s3":   {Interface: "s3", Optional: true},
				"saml": {Interface: "saml"},
			},
		},
		paascharm.GoStructsData{
			PackageName:  "pkg",
			CommonPrefix: "APP_",
			Integrations: map[string]paascharm.Integration{
				"s3":   {Name: "s3", GoName: "S3", Interface: "s3", Optional: true},
				"saml": {Name: "saml", GoName: "SAML", Interface: "saml"},
			},
		},
	}, {
		"unknown integration",
		paascharm.CharmcraftYAMLConfig{
			Requires: map[string]paascharm.CharmcraftIntegration{
				"unknown": {Interface: "Unknown"},
			},
		},
		paascharm.GoStructsData{PackageName: "pkg", CommonPrefix: "APP_"},
	}}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s", tt.name)
		t.Run(testname, func(t *testing.T) {
			actual, err := paascharm.NewGoStructsData("pkg", tt.charmcraftYAML)
			if err != nil {
				t.Fatalf("Error creating Go structs data %s", err)
			}

			if diff := pretty.Diff(actual, tt.expected); len(diff) > 0 {
				t.Errorf("go structs data is not correctly generated. diff: %s\n", diff)
			}
		})
	}
}
