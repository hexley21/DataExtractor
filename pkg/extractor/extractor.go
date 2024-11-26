package extractor

type Extractor interface {
	KeyExtractor
	DataExtractor
}

type KeyExtractor interface {
	Keys(data interface{}) []string
}

type DataExtractor interface {
	Data(data interface{}, keys []string) interface{}
}

type ExtractorImpl struct{}

func New() *ExtractorImpl {
	return &ExtractorImpl{}
}

func (ex *ExtractorImpl) Keys(data interface{}) ([]string, []string) {
	var keys []string
	var viewKeys []string
	seen := make(map[string]struct{})
	ex.extractKeys(data, "", &keys, &viewKeys, seen)
	return keys, viewKeys
}

func (ex *ExtractorImpl) extractKeys(data interface{}, prefix string, keys *[]string, viewKeys *[]string, seen map[string]struct{}) {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, value := range v {
			if _, ok := seen[key]; !ok {
				seen[key] = struct{}{}
				*keys = append(*keys, key)
				*viewKeys = append(*viewKeys, prefix+key)
			}
			ex.extractKeys(value, prefix+"  ", keys, viewKeys, seen)
		}
	case []interface{}:
		for _, item := range v {
			ex.extractKeys(item, prefix, keys, viewKeys, seen)
		}
	}
}

func (ex *ExtractorImpl) Data(data interface{}, keys []string) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for _, key := range keys {
			if val, ok := v[key]; ok {
				result[key] = ex.Data(val, keys)
			}
		}

		if len(result) == 0 {
			for key, val := range v {
				result[key] = ex.Data(val, keys)
			}
		}
		return result
	case []interface{}:
		var result []interface{}
		for _, item := range v {
			result = append(result, ex.Data(item, keys))
		}
		return result
	default:
		return v
	}
}
