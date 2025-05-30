package storage

import (
	"coinflow/coinflow-server/api-gateway/internal/clients/grpc/storage/types"
	"coinflow/coinflow-server/api-gateway/internal/models"
	pb "coinflow/coinflow-server/gen/storage_service/golang"
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"context"
	"fmt"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type StorageClient struct {
	grpcCli pb.StorageClient
	grpcCfg pkgConfig.GrpcConfig
}

func NewStorageClient(grpcCfg pkgConfig.GrpcConfig) (*StorageClient, error) {
	const op = "NewStorageClient"

	addr := fmt.Sprintf("%s:%s", grpcCfg.Host, grpcCfg.Port)
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &StorageClient{
		grpcCli: pb.NewStorageClient(conn),
		grpcCfg: grpcCfg,
	}, nil
}

func (c *StorageClient) GetTransaction(ctx context.Context, userId uuid.UUID, txId uuid.UUID) (*models.Transaction, error) {
	const op = "StorageClient.GetTransactionsInPeriod"

	req := pb.GetTransactionRequest{UserId: userId.String(), TxId: txId.String()}
	resp, err := c.grpcCli.GetTransaction(ctx, &req)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return types.CreateGetTransactionResponse(resp)
}

func (c *StorageClient) GetTransactionsInPeriod(ctx context.Context, userId uuid.UUID, begin string, end string) ([]*models.Transaction, error) {
	const op = "StorageClient.GetTransactionsInPeriod"

	req := pb.GetTransactionsInPeriodRequest{UserId: userId.String(), Begin: begin, End: end}
	resp, err := c.grpcCli.GetTransactionsInPeriod(ctx, &req)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return types.CreateGetTransactionsInPeriodResponse(resp)
}

func (c *StorageClient) PostTransaction(ctx context.Context, tx *models.Transaction, withAutoCategory bool) (string, error) {
	const op = "StorageClient.PostTransaction"

	req, err := types.CreatePostTransactionRequest(tx, withAutoCategory)
	if err != nil {
		return "", err
	}

	resp, err := c.grpcCli.PostTransaction(ctx, req)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.TxId, nil
}
