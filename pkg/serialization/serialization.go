package serialization

import (
	"errors"
	"strings"

	"github.com/hexley21/data_extractor/pkg/serialization/json"
	"github.com/hexley21/data_extractor/pkg/serialization/yaml"
)

var ErrUnsuportedExtension = errors.New("unsupported file extension")

type Processor interface {
	Serializer
	Deserializer
}

type Serializer interface {
	Serialize(data interface{}) ([]byte, error)
}

type Deserializer interface {
	Deserialize(content []byte, data interface{}) error
}

func GetProcessor(fileExtension string, indent int) (Processor, error) {
	switch strings.ToLower(fileExtension) {
	case ".json":
		return json.New(indent), nil
	case ".yaml", ".yml":
		return yaml.New(indent), nil
	}
	return nil, ErrUnsuportedExtension
}
