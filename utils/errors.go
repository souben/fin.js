package utils

import "fmt"

type SanitizingError struct {
	Value interface{}
	Err   string
}

func (serr *SanitizingError) Error() string {
	return fmt.Sprintf("SanitizingError for %v: %v", serr.Value, serr.Err)
}
