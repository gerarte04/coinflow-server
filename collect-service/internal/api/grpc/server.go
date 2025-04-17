package grpc

import (
	"coinflow/coinflow-server/collect-service/internal/api/grpc/types"
	"coinflow/coinflow-server/collect-service/internal/usecases"
	pb "coinflow/coinflow-server/gen/collect_service"
	"context"
)

type CollectServer struct {
	pb.UnimplementedCollectServer
	collectSvc usecases.CollectService
}

func NewCoinflowServer(collectSvc usecases.CollectService) *CollectServer {
	return &CollectServer{
		collectSvc: collectSvc,
	}
}

func (s *CollectServer) GetTransactionCategory(ctx context.Context, r *pb.GetTransactionCategoryRequest) (*pb.GetTransactionCategoryResponse, error) {
	reqObj, err := types.CreateGetTransactionCategoryRequestObject(r)
	if err != nil {
		return nil, CreateRequestObjectStatusError(err)
	}

	err = s.collectSvc.CollectCategory(reqObj.Ts)
	if err != nil {
	    return nil, CreateResultStatusError(err)
	}

	resp, err := types.CreateGetTransactionCategoryResponse("category")
	if err != nil {
		return nil, CreateResponseStatusError(err)
	}

	return resp, nil
}
