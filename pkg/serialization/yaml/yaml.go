package yaml

import (
	"bytes"

	"github.com/hexley21/data_extractor/pkg/config"
	"gopkg.in/yaml.v3"
)

type YamlProcessor struct{
	beautify config.YamlBeautify
}

func New(beautify *config.YamlBeautify) *YamlProcessor {
	if beautify == nil {
		return &YamlProcessor{config.YamlBeautify{}}
	}
	return &YamlProcessor{*beautify}
}

func (p *YamlProcessor) Serialize(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	e := yaml.NewEncoder(&buf)
	e.SetIndent(p.beautify.Indent)
	defer e.Close()

	if err := e.Encode(data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (p *YamlProcessor) Deserialize(content []byte, data interface{}) error {
	return yaml.Unmarshal(content, data)
}
