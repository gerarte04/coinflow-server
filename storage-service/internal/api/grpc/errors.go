package grpc

import (
	"coinflow/coinflow-server/pkg/utils"
	"coinflow/coinflow-server/storage-service/internal/repository"
	"errors"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var(
	errorCodes = map[error]codes.Code{
		utils.ErrorInvalidId: codes.InvalidArgument,

		repository.ErrorTxIdNotFound: codes.NotFound,
		repository.ErrorUserIdNotFound: codes.NotFound,
		repository.ErrorNoSuchCredExists: codes.PermissionDenied,
	}
)

func CreateRequestObjectStatusError(err error) error {
	log.Printf("%s\n", err.Error())

	return status.Error(codes.InvalidArgument, fmt.Sprintf("creating request object: %s", err.Error()))
}

func CreateResultStatusError(err error) error {
	log.Printf("%s\n", err.Error())

	var basicErr error
	for next := err; next != nil; next = errors.Unwrap(basicErr) {
		basicErr = next
	}

	code, ok := errorCodes[basicErr]
	if !ok {
		code = codes.Internal
	}

	return status.Error(code, err.Error())
}

func CreateResponseStatusError(err error) error {
	log.Printf("%s\n", err.Error())

	return status.Error(codes.Internal, fmt.Sprintf("creating response: %s", err.Error()))
}
