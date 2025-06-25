package grpc

import (
	pb "coinflow/coinflow-server/gen/storage_service/golang"
	pkgGrpc "coinflow/coinflow-server/pkg/grpc"
	grpcErr "coinflow/coinflow-server/pkg/pkgerrors/grpc"
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

func (s *StorageServer) GetTransaction(ctx context.Context, r *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
	reqObj, err := types.CreateGetTransactionRequestObject(r)
	if err != nil {
		return nil, grpcErr.CreateRequestObjectStatusError(err)
	}

	tx, err := s.txService.GetTransaction(ctx, reqObj.UserId, reqObj.TxId)
	if err != nil {
	    return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	resp, err := types.CreateGetTransactionResponse(tx)
	if err != nil {
		return nil, grpcErr.CreateResponseStatusError(err)
	}

	return resp, nil
}

func (s *StorageServer) GetTransactionsInPeriod(ctx context.Context, r *pb.GetTransactionsInPeriodRequest) (*pb.GetTransactionsInPeriodResponse, error) {
	reqObj, err := types.CreateGetTransactionsInPeriodRequestObject(r)
	if err != nil {
		return nil, grpcErr.CreateRequestObjectStatusError(err)
	}

	txs, err := s.txService.GetTransactionsInPeriod(ctx, reqObj.UserId, reqObj.Begin, reqObj.End)
	if err != nil {
	    return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	resp, err := types.CreateGetTransactionsInPeriodResponse(txs)
	if err != nil {
		return nil, grpcErr.CreateResponseStatusError(err)
	}

	return resp, nil
}

func (s *StorageServer) PostTransaction(ctx context.Context, r *pb.PostTransactionRequest) (*pb.PostTransactionResponse, error) {
	reqObj, err := types.CreatePostTransactionRequestObject(r)
	if err != nil {
		return nil, grpcErr.CreateRequestObjectStatusError(err)
	}

	txId, err := s.txService.PostTransaction(ctx, reqObj.Tx, reqObj.WithAutoCategory)
	if err != nil {
	    return nil, grpcErr.CreateResultStatusError(err, errorCodes)
	}

	pkgGrpc.SetResponseCode(ctx, 201)

	return &pb.PostTransactionResponse{TxId: txId.String()}, nil
}
