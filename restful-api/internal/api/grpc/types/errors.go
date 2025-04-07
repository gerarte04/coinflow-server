package types

import "errors"

var (
    ErrorEmptyId = errors.New("Couldn't extract id from request")
)
