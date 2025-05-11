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

func (c *StorageClient) GetTransaction(txId uuid.UUID) (*models.Transaction, error) {
	const op = "StorageClient.GetTransactionsInPeriod"

	req := pb.GetTransactionRequest{TxId: txId.String()}
	resp, err := c.grpcCli.GetTransaction(context.Background(), &req)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return types.CreateGetTransactionResponse(resp)
}

func (c *StorageClient) GetTransactionsInPeriod(begin string, end string) ([]*models.Transaction, error) {
	const op = "StorageClient.GetTransactionsInPeriod"

	req := pb.GetTransactionsInPeriodRequest{Begin: begin, End: end}
	resp, err := c.grpcCli.GetTransactionsInPeriod(context.Background(), &req)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return types.CreateGetTransactionsInPeriodResponse(resp)
}

func (c *StorageClient) PostTransaction(tx *models.Transaction) (string, error) {
	const op = "StorageClient.PostTransaction"

	req, err := types.CreatePostTransactionRequest(tx)
	if err != nil {
		return "", err
	}

	resp, err := c.grpcCli.PostTransaction(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return resp.TxId, nil
}
