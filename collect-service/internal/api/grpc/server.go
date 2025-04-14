package grpc

import (
	"coinflow/coinflow-server/collect-service/internal/api/grpc/types"
	pb "coinflow/coinflow-server/gen/collect_service"
	"context"
)

type CollectServer struct {
	pb.UnimplementedCollectServer
}

func NewCoinflowServer() *CollectServer {
	return &CollectServer{}
}

func (s *CollectServer) GetTransactionCategory(ctx context.Context, r *pb.GetTransactionCategoryRequest) (*pb.GetTransactionCategoryResponse, error) {
	_, err := types.CreateGetTransactionCategoryRequestObject(r)
	if err != nil {
		return nil, CreateRequestObjectStatusError(err)
	}

	// res, err := s.tsService.GetTransaction(reqObj.TsId)
	// if err != nil {
	//     return nil, CreateResultStatusError(err)
	// }

	resp, err := types.CreateGetTransactionCategoryResponse("category")
	if err != nil {
		return nil, CreateResponseStatusError(err)
	}

	return resp, nil
}
