package serialization

import (
	"errors"
	"strings"

	"github.com/hexley21/data_extractor/pkg/config"
	"github.com/hexley21/data_extractor/pkg/serialization/json"
)

var ErrUnsuportedExtension = errors.New("unsupported file extension")

func GetProcessor(fileExtension string, cfg config.Beautify, beautify bool) (Processor, error) {
	switch strings.ToLower(fileExtension) {
	case ".json":
		if beautify {
			return json.New(cfg.Prefix, cfg.Indent), nil
		}
		return json.New("", ""), nil
	default:
		return nil, ErrUnsuportedExtension
	}
}
