package grpc

import (
	"coinflow/coinflow-server/collection-service/internal/usecases"
	pb "coinflow/coinflow-server/gen/collection_service/golang"
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
