package grpc

import (
	"errors"

	"google.golang.org/grpc/codes"
)

var (
	ErrorUserIdsDontMatch = errors.New("user ids in request and JWT aren't matching")

	errorCodes = map[error]codes.Code{
		ErrorUserIdsDontMatch: codes.PermissionDenied,
	}
)
