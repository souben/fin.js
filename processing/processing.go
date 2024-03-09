package processing

import (
	"strings"

	"github.com/souben/fin.js/utils"
)

var (
	trueBool    = []string{"1", "t", "T", "TRUE", "true", "True"}
	falseBool   = []string{"0", "f", "F", "FALSE", "false", "False"}
	emptyString = ""
)

func TransformAndProcess(input map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range input {
		key, err := sanitize(key)
		if err != nil {
			continue
		}
		value, ok := value.(map[string]interface{})
		if !ok {
			continue
		}
		a, err := transform(value)
		if err != nil {
			continue
		}
		result[key] = a
	}
	return result
}

func transform(data map[string]interface{}) (interface{}, error) {
	for key, value := range data {
		key, err := sanitize(key) // check earlier if key  is empty string
		if err != nil {
			return nil, err
		}
		return process(key, value)
	}
	return nil, nil
}

func process(key string, value interface{}) (interface{}, error) {
	if isSimpleType(value) {
		return simpleProcessing(key, value.(string))
	}
	return complexProcessing(key, value)
}

func sanitize(value interface{}) (string, error) {
	if !isSimpleType(value) {
		return emptyString, &utils.SanitizingError{Value: value, Err: "Unsupported type"}
	}
	s := strings.Trim(value.(string), " ")
	if s == emptyString {
		return emptyString, &utils.SanitizingError{Value: s, Err: "Empty string"}
	}
	return s, nil
}

func isSimpleType(value interface{}) bool {
	_, ok := value.(string)
	return ok
}
