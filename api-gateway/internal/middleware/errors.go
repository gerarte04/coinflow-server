package middleware

import "errors"

var (
	ErrorTokenNotFound = errors.New("auth middleware: couldn't extract token neither from header or cookie")
)
