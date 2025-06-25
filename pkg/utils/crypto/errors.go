package crypto

import "errors"

var (
	ErrorTokenParsingFailed = errors.New("token parsing failed, check if it's correct")
	ErrorTokenSignatureInvalid = errors.New("token signature is invalid")
	ErrorUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrorTokenExpired = errors.New("Token has been expired")
)
