package config

import "gopkg.in/yaml.v3"

type (
	Config struct {
		AppName string `yaml:"app_name"`
		CliName string `yaml:"cli_name"`

		Colors struct {
			MultiSelect map[string]MultiSelect `yaml:"multi_select"`
		}
	}

	MultiSelect struct {
		Focused  string `yaml:"focused"`
		Selected string `yaml:"selected"`
	}
)

func LoadConfig(file []byte) (*Config, error) {
	cfg, err := parseEmbed(file)
	if err != nil {
		return nil, err
	}

	cfg.Colors.MultiSelect[".yaml"] = cfg.Colors.MultiSelect[".yml"]
	return cfg, nil
}

func parseEmbed(file []byte) (*Config, error) {
	config := &Config{}
	err := yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
