package utils

import "errors"

var (
	ErrorInvalidId = errors.New("couldn't extract id from request")
)
