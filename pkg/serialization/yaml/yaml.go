package yaml

import (
	"bytes"

	"gopkg.in/yaml.v3"
)

type YamlProcessor struct{
	indent int
}

func New() *YamlProcessor {
	return &YamlProcessor{4}
}

func (p *YamlProcessor) Serialize(data interface{}) ([]byte, error) {
	var buf bytes.Buffer
	e := yaml.NewEncoder(&buf)
	e.SetIndent(p.indent)
	defer e.Close()

	if err := e.Encode(data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (p *YamlProcessor) Deserialize(content []byte, data interface{}) error {
	return yaml.Unmarshal(content, data)
}
