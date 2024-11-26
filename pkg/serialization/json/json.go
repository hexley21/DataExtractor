package json

import (
	"encoding/json"

	"github.com/hexley21/data_extractor/pkg/config"
)

type JsonProcessor struct{
	beautify *config.JsonBeutify
}

func New(beautify *config.JsonBeutify) *JsonProcessor {
	return &JsonProcessor{beautify}
}

func (p *JsonProcessor) Serialize(data interface{}) ([]byte, error) {
	if p.beautify == nil {
		return json.Marshal(data)
	}
	return json.MarshalIndent(data, p.beautify.Prefix, p.beautify.Indent)
}

func (p *JsonProcessor) Deserialize(content []byte, data interface{}) error {
	return json.Unmarshal(content, data)
}
