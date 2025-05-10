package grpc

import (
	"coinflow/coinflow-server/pkg/pkgerrors"
	"fmt"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateRequestObjectStatusError(err error) error {
	log.Printf("%s\n", err.Error())

	return status.Error(codes.InvalidArgument, fmt.Sprintf("creating request object: %s", err.Error()))
}

func CreateResultStatusError(err error, errorCodes map[error]codes.Code) error {
	log.Printf("%s\n", err.Error())

	basicErr := pkgerrors.UnwrapAll(err)
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
