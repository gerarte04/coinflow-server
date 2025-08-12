package types

import "errors"

var (
	ErrorEmptyId = errors.New("Couldn't extract id from request")
	ErrorInvalidPageSize = errors.New("invalid page size, must be non-negative")
	ErrorInvalidPageToken = errors.New("invalid page token")
)
