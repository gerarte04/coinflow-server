package grpc

import (
	pb "coinflow/coinflow-server/gen/cfapi"
	"coinflow/coinflow-server/restful-api/internal/api/grpc/types"
	"coinflow/coinflow-server/restful-api/internal/usecases"
	"context"
)

type CoinflowServer struct {
    pb.UnimplementedCoinflowServer
    tsService usecases.TransactionsService
}

func NewCoinflowServer(tsService usecases.TransactionsService) *CoinflowServer {
    return &CoinflowServer{tsService: tsService}
}

func (s *CoinflowServer) GetTransaction(ctx context.Context, r *pb.GetTransactionRequest) (*pb.GetTransactionResponse, error) {
    reqObj, err := types.CreateGetTransactionRequestObject(r)
    if err != nil {
        return nil, CreateRequestObjectStatusError(err)
    }

    res, err := s.tsService.GetTransaction(reqObj.TsId)
    if err != nil {
        return nil, CreateResultStatusError(err)
    }

    resp, err := types.CreateGetTransactionResponse(res)
    if err != nil {
        return nil, CreateResponseStatusError(err)
    }

    return resp, nil
}

func (s *CoinflowServer) PostTransaction(ctx context.Context, r *pb.PostTransactionRequest) (*pb.PostTransactionResponse, error) {
    reqObj, err := types.CreatePostTransactionRequestObject(r)
    if err != nil {
        return nil, CreateRequestObjectStatusError(err)
    }

    res, err := s.tsService.PostTransaction(reqObj.Ts)
    if err != nil {
        return nil, CreateResultStatusError(err)
    }

    return &pb.PostTransactionResponse{TsId: res}, nil
}
