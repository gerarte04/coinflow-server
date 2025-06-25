package grpc

import (
	"context"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func SetResponseCode(ctx context.Context, code int) {
	grpc.SetHeader(ctx, metadata.Pairs("x-http-code", strconv.Itoa(code)))
}
