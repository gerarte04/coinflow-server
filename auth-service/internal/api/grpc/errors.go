package grpc

import (
	"coinflow/coinflow-server/auth-service/internal/repository"
	"coinflow/coinflow-server/auth-service/internal/usecases"
	"coinflow/coinflow-server/pkg/utils"
	"coinflow/coinflow-server/pkg/utils/crypto"

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

		crypto.ErrorTokenExpired: codes.Unauthenticated,
		crypto.ErrorTokenParsingFailed: codes.Unauthenticated,
		crypto.ErrorTokenSignatureInvalid: codes.Unauthenticated,
		crypto.ErrorUnexpectedSigningMethod: codes.Unauthenticated,
	}
)
