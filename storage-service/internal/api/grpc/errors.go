package grpc

import (
	"coinflow/coinflow-server/pkg/utils"
	"coinflow/coinflow-server/storage-service/internal/repository"
	"errors"

	"google.golang.org/grpc/codes"
)

var(
	ErrorUserIdsDontMatch = errors.New("user ids in request and header aren't matching")

	errorCodes = map[error]codes.Code{
		utils.ErrorInvalidId: codes.InvalidArgument,

		repository.ErrorTxIdNotFound: codes.NotFound,
		repository.ErrorPermissionDenied: codes.PermissionDenied,

		ErrorUserIdsDontMatch: codes.PermissionDenied,
	}
)
