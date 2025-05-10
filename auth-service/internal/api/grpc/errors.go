package grpc

import (
	"coinflow/coinflow-server/auth-service/internal/repository"
	"coinflow/coinflow-server/auth-service/internal/usecases"
	"coinflow/coinflow-server/pkg/utils"

	"google.golang.org/grpc/codes"
)

var(
	errorCodes = map[error]codes.Code{
		utils.ErrorInvalidId: codes.InvalidArgument,

		repository.ErrorUserIdNotFound: codes.NotFound,
		repository.ErrorUserCredAlreadyExists: codes.Unauthenticated,
		repository.ErrorUserLoginNotFound: codes.Unauthenticated,
		repository.ErrorWrongPassword: codes.Unauthenticated,

		usecases.ErrorTokenInBlacklist: codes.Unauthenticated,
	}
)
