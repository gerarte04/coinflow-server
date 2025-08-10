package grpc

import (
	pb "coinflow/coinflow-server/gen/storage_service/golang"
	pkgGrpc "coinflow/coinflow-server/pkg/grpc"
	grpcErr "coinflow/coinflow-server/pkg/pkgerrors/grpc"
	"coinflow/coinflow-server/storage-service/config"
	"coinflow/coinflow-server/storage-service/internal/api/grpc/types"
	"coinflow/coinflow-server/storage-service/internal/usecases"
	"context"
	"fmt"
)

type StorageServer struct {
	pb.UnimplementedStorageServer
	txService usecases.TransactionsService
	cfg config.ServiceConfig
}

func NewStorageServer(txService usecases.TransactionsService, cfg config.ServiceConfig) *StorageServer {
	return &StorageServer{
		txService: txService,
		cfg: cfg,
	}
}

func (s *StorageServer) GetTransaction(ctx context.Context, r *pb.GetTransactionRequest) (*pb.Transaction, error) {
	reqObj, err := types.CreateGetTransactionRequestObject(ctx, r)
	if err != nil {
		return nil, grpcErr.CreateRequestObjectStatusError(err)
	}

	tx, err := s.txService.GetTransaction(ctx, reqObj.UserId, reqObj.TxId)
	if err != nil {
	    return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	resp, err := types.GetProtobufTxFromModel(tx)
	if err != nil {
		return nil, grpcErr.CreateResponseStatusError(err)
	}

	return resp, nil
}

func (s *StorageServer) ListTransactions(ctx context.Context, r *pb.ListTransactionsRequest) (*pb.ListTransactionsResponse, error) {
	reqObj, err := types.CreateListTransactionsRequestObject(ctx, r)
	if err != nil {
		return nil, grpcErr.CreateRequestObjectStatusError(err)
	}

	if reqObj.UserId.String() != r.UserId {
		return nil, grpcErr.CreateResultStatusError(
			fmt.Errorf("ListTransactions: %w", ErrorUserIdsDontMatch),
			errorCodes,
		)
	}

	txs, err := s.txService.GetTransactionsInPeriod(ctx, reqObj.UserId, reqObj.Begin, reqObj.End)
	if err != nil {
	    return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	resp, err := types.CreateListTransactionsResponse(txs)
	if err != nil {
		return nil, grpcErr.CreateResponseStatusError(err)
	}

	return resp, nil
}

func (s *StorageServer) CreateTransaction(ctx context.Context, r *pb.CreateTransactionRequest) (*pb.Transaction, error) {
	reqObj, err := types.MakeCreateTransactionRequestObject(ctx, r)
	if err != nil {
		return nil, grpcErr.CreateRequestObjectStatusError(err)
	}

	if reqObj.Tx.UserId.String() != r.UserId {
		return nil, grpcErr.CreateResultStatusError(
			fmt.Errorf("CreateTransaction: %w", ErrorUserIdsDontMatch),
			errorCodes,
		)
	}

	tx, err := s.txService.PostTransaction(ctx, reqObj.Tx, reqObj.WithAutoCategory)
	if err != nil {
	    return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	pkgGrpc.SetResponseCode(ctx, s.cfg.HttpCodeHeaderName, 201)

	resp, err := types.GetProtobufTxFromModel(tx)
	if err != nil {
		return nil, grpcErr.CreateResponseStatusError(err)
	}

	return resp, nil
}
