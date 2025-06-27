package grpc

import (
	"context"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func GetHeader(ctx context.Context, name string) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("couldn't retrieve metadata from context")
	}

	value := md.Get(name)
	if len(value) == 0 {
		return "", fmt.Errorf("no header named %s", name)
	}

	return value[0], nil
}

func SetResponseCode(ctx context.Context, name string, code int) {
	grpc.SetHeader(ctx, metadata.Pairs(name, strconv.Itoa(code)))
}
