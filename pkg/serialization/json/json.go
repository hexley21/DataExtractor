package json

import (
	jsoniter "github.com/json-iterator/go"
	"strings"
)

type JsonProcessor struct {
	indent string
}

func New(indent int) *JsonProcessor {
	return &JsonProcessor{strings.Repeat(" ", indent)}
}

func (p *JsonProcessor) Serialize(data interface{}) ([]byte, error) {
	if p.indent == "" {
		return jsoniter.Marshal(data)
	}
	return jsoniter.MarshalIndent(data, "", p.indent)
}

func (p *JsonProcessor) Deserialize(content []byte, data interface{}) error {
	return jsoniter.Unmarshal(content, data)
}
