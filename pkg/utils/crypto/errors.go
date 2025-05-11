package crypto

import "errors"

var (
	ErrorUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrorTokenExpired = errors.New("Token has been expired")
)
