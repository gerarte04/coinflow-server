package grpc

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func CreateRequestObjectStatusError(err error) error {
	return status.Error(codes.InvalidArgument, fmt.Sprintf("creating request object: %s", err.Error()))
}

func CreateResultStatusError(err error) error {
	var code codes.Code
	var msg string

	switch err {
	default:
		code = codes.Internal
		msg = fmt.Sprintf("unexpected error: %s", err.Error())
	}

	return status.Error(code, msg)
}

func CreateResponseStatusError(err error) error {
	return status.Error(codes.Internal, fmt.Sprintf("creating response: %s", err.Error()))
}
