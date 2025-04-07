package types

import "errors"

var (
    ErrorInvalidId = errors.New("couldn't extract id from request")
    ErrorParseTransaction = errors.New("couldn't parse transaction from request body")
)
