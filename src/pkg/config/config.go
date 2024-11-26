package config

import (
	"gopkg.in/yaml.v3"
)

type (
	Config struct {
		AppName  string              `yaml:"app_name"`
		CliName  string              `yaml:"cli_name"`
		Beautify map[string]Beautify `yaml:"beautify"`
		Colors   struct {
			MultiSelect map[string]MultiSelect `yaml:"multi_select"`
		}
	}

	MultiSelect struct {
		Focused  string `yaml:"focused"`
		Selected string `yaml:"selected"`
	}

	Beautify struct {
		Prefix string `yaml:"prefix"`
		Indent string `yaml:"indent"`
	}
)

func LoadConfig(file []byte) (*Config, error) {
	return parseEmbed(file)
}

func parseEmbed(file []byte) (*Config, error) {
	config := &Config{}
	err := yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
