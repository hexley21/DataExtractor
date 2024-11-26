package json

import (
	"encoding/json"
	"strings"
)

type JsonProcessor struct{
	indent string
}

func New(indent int) *JsonProcessor {
	if (indent % 4) == 0 {
		return &JsonProcessor{strings.Repeat("\t", indent/4)}
	} 
	return &JsonProcessor{strings.Repeat(" ", indent)}
}

func (p *JsonProcessor) Serialize(data interface{}) ([]byte, error) {
	if p.indent == "" {
		return json.Marshal(data)
	}
	return json.MarshalIndent(data, "", p.indent)
}

func (p *JsonProcessor) Deserialize(content []byte, data interface{}) error {
	return json.Unmarshal(content, data)
}
