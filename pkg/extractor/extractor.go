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

func (ex *ExtractorImpl) Keys(data interface{}) []string {
	switch data := data.(type) {
	case map[string]interface{}:
		return ex.objectKeys(data)
	case []interface{}:
		return ex.arrayKeys(data, &[]string{})
	default:
		return nil
	}
}

func (ex *ExtractorImpl) objectKeys(data map[string]interface{}) []string {
	keys := make([]string, len(data))

	var i int
	for key := range data {
		keys[i] = key
		i++
	}

	return keys
}

func (ex *ExtractorImpl) arrayKeys(array []interface{}, keys *[]string) []string {
	for _, data := range array {
		switch data := data.(type) {
		case map[string]interface{}:
			*keys = append(*keys, ex.objectKeys(data)...)
		case []interface{}:
			return ex.arrayKeys(data, keys)
		default:
			return nil
		}
	}

	return removeDuplicateKeys(*keys)
}

func removeDuplicateKeys(slice []string) []string {
	seen := make(map[string]bool)
	result := []string{}

	for _, value := range slice {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}

	return result
}

func (ex *ExtractorImpl) Data(data interface{}, keys []string) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		return ex.objectData(v, keys)
	case []interface{}:
		return ex.arrayData(v, keys)
	default:
		return v
	}
}

func (ex *ExtractorImpl) arrayData(data []interface{}, keys []string) interface{} {
	var result []interface{}
	for _, item := range data {
		result = append(result, ex.Data(item, keys))
	}
	return result
}

func (ex *ExtractorImpl) objectData(data map[string]interface{}, keys []string) interface{} {
	result := make(map[string]interface{})
	for _, key := range keys {
		if val, ok := data[key]; ok {
			result[key] = ex.Data(val, keys)
		}
	}

	if len(result) > 0 {
		return result
	}

	for key, val := range data {
		result[key] = ex.Data(val, keys)
	}
	return result
}
