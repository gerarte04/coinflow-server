package types

import "errors"

var (
	ErrorParseTransaction = errors.New("couldn't parse transaction from request body")
)
