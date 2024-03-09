package processing

import (
	"errors"
	"reflect"
)

func complexProcessing(key string, value interface{}) (interface{}, error) {
	switch key {
	case "L":
		return listProcessing(value)
	case "M":
		return mapProcessing(value)
	default:
		return nil, errors.New("invalid complex type")
	}
}

func listProcessing(list interface{}) ([]interface{}, error) {
	result := make([]interface{}, 0)
	if reflect.ValueOf(list).Kind() == reflect.Slice {
		for _, item := range list.([]interface{}) {
			a, ok := item.(map[string]interface{})
			if !ok || len(a) != 1 {
				return nil, errors.New("invalid L type")
			}
			for key, value := range a {
				if !isSimpleType(value) {
					continue
				}

				value, err := sanitize(value)
				if err != nil {
					continue
				}

				itemData, err := simpleProcessing(key, value)
				if err != nil {
					continue
				}
				result = append(result, itemData)
			}
		}
		return result, nil
	}
	return nil, errors.New("invalid L type")
}

func mapProcessing(mMap interface{}) (map[string]interface{}, error) {
	var result map[string]interface{}
	if reflect.ValueOf(mMap).Kind() == reflect.Map {
		v, ok := mMap.(map[string]interface{})
		if ok {
			result = TransformAndProcess(v)
			return result, nil
		}
	}
	return nil, errors.New("invalid M type")
}
