package json

import (
	"encoding/json"
)

type JsonProcessor struct{
	prefix string
	indent string
}

func New(prefix string, indent string) *JsonProcessor {
	return &JsonProcessor{prefix, indent}
}

func (p *JsonProcessor) Serialize(data interface{}) ([]byte, error) {
	if p.prefix == "" && p.indent == "" {
		return json.Marshal(data)
	}
	return json.MarshalIndent(data, p.prefix, p.indent)
}

func (p *JsonProcessor) Deserialize(content []byte, data interface{}) error {
	return json.Unmarshal(content, &data)
}
