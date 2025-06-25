package grpc

import (
	"coinflow/coinflow-server/pkg/utils"
	"coinflow/coinflow-server/storage-service/internal/repository"

	"google.golang.org/grpc/codes"
)

var(
	errorCodes = map[error]codes.Code{
		utils.ErrorInvalidId: codes.InvalidArgument,

		repository.ErrorTxIdNotFound: codes.NotFound,
		repository.ErrorPermissionDenied: codes.PermissionDenied,
	}
)
