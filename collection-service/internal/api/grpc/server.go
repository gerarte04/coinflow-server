package grpc

import (
	"coinflow/coinflow-server/collection-service/internal/api/grpc/types"
	"coinflow/coinflow-server/collection-service/internal/usecases"
	pb "coinflow/coinflow-server/gen/collection_service/golang"
	pkgGrpc "coinflow/coinflow-server/pkg/pkgerrors/grpc"
	"context"
	"fmt"
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

func (s *CollectionServer) GetSummaryInPeriod(
	ctx context.Context,
	r *pb.GetSummaryInPeriodRequest,
) (*pb.GetSummaryInPeriodResponse, error) {
	req, err := types.CreateGetSummaryInPeriodRequestObject(ctx, r)
	if err != nil {
		return nil, pkgGrpc.CreateRequestObjectStatusError(err)
	}

	if req.UserId.String() != r.UserId {
		return nil, pkgGrpc.CreateResultStatusError(
			fmt.Errorf("%w", ErrorUserIdsDontMatch),
			errorCodes,
		)
	}

	res, err := s.collectSvc.GetSummaryInPeriod(ctx, req.UserId, req.Begin, req.End)
	if err != nil {
		return nil, pkgGrpc.CreateResultStatusError(err, errorCodes)
	}

	resp, err := types.GetModelSummaryFromProtobuf(res)
	if err != nil {
		return nil, pkgGrpc.CreateResponseStatusError(err)
	}

	return &pb.GetSummaryInPeriodResponse{Summary: resp}, nil
}

func (s *CollectionServer) GetSummaryInLastNMonths(
	ctx context.Context,
	r *pb.GetSummaryInLastNMonthsRequest,
) (*pb.GetSummaryInLastNMonthsResponse, error) {
	req, err := types.CreateGetSummaryInLastNMonthsRequestObject(ctx, r)
	if err != nil {
		return nil, pkgGrpc.CreateRequestObjectStatusError(err)
	}

	if req.UserId.String() != r.UserId {
		return nil, pkgGrpc.CreateResultStatusError(
			fmt.Errorf("%w", ErrorUserIdsDontMatch),
			errorCodes,
		)
	}

	res, err := s.collectSvc.GetSummaryInLastNMonths(ctx, req.UserId, req.N, req.CurTime, req.Timezone)
	if err != nil {
		return nil, pkgGrpc.CreateResultStatusError(err, errorCodes)
	}

	resp, err := types.CreateGetSummaryInLastNMonthsResponse(res)
	if err != nil {
		return nil, pkgGrpc.CreateResponseStatusError(err)
	}

	return resp, nil
}

func (s CollectionServer) GetSummaryByCategories(
	ctx context.Context,
	r *pb.GetSummaryByCategoriesRequest,
) (*pb.GetSummaryByCategoriesResponse, error) {
	req, err := types.CreateGetSummaryByCategoriesRequestObject(ctx, r)
	if err != nil {
		return nil, pkgGrpc.CreateRequestObjectStatusError(err)
	}

	if req.UserId.String() != r.UserId {
		return nil, pkgGrpc.CreateResultStatusError(
			fmt.Errorf("%w", ErrorUserIdsDontMatch),
			errorCodes,
		)
	}

	res, err := s.collectSvc.GetSummaryByCategories(ctx, req.UserId, req.Begin, req.End)
	if err != nil {
		return nil, pkgGrpc.CreateResultStatusError(err, errorCodes)
	}

	resp, err := types.CreateGetSummaryByCategoriesResponse(res)
	if err != nil {
		return nil, pkgGrpc.CreateResponseStatusError(err)
	}

	return resp, nil
}
