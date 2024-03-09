package processing

import (
	"errors"
	"slices"
	"strconv"
	"time"
)

func simpleProcessing(key string, value string) (interface{}, error) {
	switch key {
	case "S":
		return StringProcessing(value)
	case "N":
		return NumericProcessing(value)
	case "BOOL":
		return BooleanProcessing(value)
	case "NULL":
		return NullProcessing(value)
	default:
		return nil, errors.New("invalid simple type")
	}
}

func StringProcessing(s string) (interface{}, error) {
	if data, err := time.Parse(time.RFC3339, s); err == nil {
		return float64(data.Unix()), nil
	}

	s, err := sanitize(s)
	if err != nil {
		return emptyString, nil
	}
	return s, nil
}

func NumericProcessing(s string) (float64, error) {
	if data, err := strconv.ParseInt(s, 10, 64); err == nil {
		return float64(data), nil
	}
	if data, err := strconv.ParseFloat(s, 64); err == nil {
		return data, nil
	}
	return 0, errors.New("invalid N type")
}

func BooleanProcessing(s string) (bool, error) {
	if slices.Contains(trueBool, s) {
		return true, nil
	}

	if slices.Contains(falseBool, s) {
		return false, nil
	}
	return false, errors.New("invalid BOOL type")

}

func NullProcessing(s string) (bool, error) {
	if isNull, _ := BooleanProcessing(s); isNull {
		return isNull, nil
	}
	return false, errors.New("invalid NULL type")
}
