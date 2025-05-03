package grpc

import (
	pb "coinflow/coinflow-server/gen/storage_service/golang"
	"coinflow/coinflow-server/storage-service/internal/api/grpc/types"
	"coinflow/coinflow-server/storage-service/internal/usecases"
	"context"
)

type StorageServer struct {
	pb.UnimplementedStorageServer
	txService usecases.TransactionsService
}

func NewStorageServer(txService usecases.TransactionsService) *StorageServer {
	return &StorageServer{
		txService: txService,
	}
}

func (s *StorageServer) GetTransactionsInPeriod(ctx context.Context, r *pb.GetTransactionsInPeriodRequest) (*pb.GetTransactionsInPeriodResponse, error) {
	reqObj, err := types.CreateGetTransactionsInPeriodRequestObject(r)
	if err != nil {
		return nil, CreateRequestObjectStatusError(err)
	}

	txs, err := s.txService.GetTransactionsInPeriod(reqObj.Begin, reqObj.End)
	if err != nil {
	    return nil, CreateResultStatusError(err)
	}

	resp, err := types.CreateGetTransactionsInPeriodResponse(txs)
	if err != nil {
		return nil, CreateResponseStatusError(err)
	}

	return resp, nil
}

func (s *StorageServer) PostTransaction(ctx context.Context, r *pb.PostTransactionRequest) (*pb.PostTransactionResponse, error) {
	reqObj, err := types.CreatePostTransactionRequestObject(r)
	if err != nil {
		return nil, CreateRequestObjectStatusError(err)
	}

	txId, err := s.txService.PostTransaction(reqObj.Tx)
	if err != nil {
	    return nil, CreateResultStatusError(err)
	}

	resp, err := types.CreatePostTransactionResponse(txId)
	if err != nil {
		return nil, CreateResponseStatusError(err)
	}

	return resp, nil
}
