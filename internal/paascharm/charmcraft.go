package paascharm

import (
	"io"

	"gopkg.in/yaml.v3"
)

func ParseCharmcraftYAML(reader io.Reader) (CharmcraftYAMLConfig, error) {
	decoder := yaml.NewDecoder(reader)
	var charmcraft CharmcraftYAMLConfig
	err := decoder.Decode(&charmcraft)
	if err != nil {
		return CharmcraftYAMLConfig{}, err
	}
	return charmcraft, nil
}

type CharmcraftYAMLConfig struct {
	Config   CharmcraftConfig                 `yaml:"config"`
	Requires map[string]CharmcraftIntegration `yaml:"requires"`
}

type CharmcraftConfig struct {
	Options map[string]CharmcraftConfigOption `yaml:"options"`
}

type CharmcraftConfigOption struct {
	Type        string `yaml:"type"`
	Default     any    `yaml:"default"`
	Description string `yaml:"description"`
}

type CharmcraftIntegration struct {
	Interface string `yaml:"interface"`
	Optional  bool   `yaml:"optional"`
}
