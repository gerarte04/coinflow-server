package grpc

import (
	"coinflow/coinflow-server/collection-service/internal/api/grpc/types"
	"coinflow/coinflow-server/collection-service/internal/usecases"
	pb "coinflow/coinflow-server/gen/collection_service/golang"
	"context"
)

type CollectionServer struct {
	pb.UnimplementedCollectionServer
	collectSvc usecases.CollectionService
}

func NewCollectionServer(collectSvc usecases.CollectionService) *CollectionServer {
	return &CollectionServer{
		collectSvc: collectSvc,
	}
}

func (s *CollectionServer) GetTransactionCategory(ctx context.Context, r *pb.GetTransactionCategoryRequest) (*pb.GetTransactionCategoryResponse, error) {
	reqObj, err := types.CreateGetTransactionCategoryRequestObject(r)
	if err != nil {
		return nil, CreateRequestObjectStatusError(err)
	}

	cat, err := s.collectSvc.CollectCategory(ctx, reqObj.Tx)
	if err != nil {
	    return nil, CreateResultStatusError(err)
	}

	resp, err := types.CreateGetTransactionCategoryResponse(cat)
	if err != nil {
		return nil, CreateResponseStatusError(err)
	}

	return resp, nil
}
