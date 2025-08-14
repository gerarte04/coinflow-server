package types

import (
	"context"
	"fmt"
	"time"

	"coinflow/coinflow-server/collection-service/internal/models"
	pb "coinflow/coinflow-server/gen/collection_service/golang"
	pkgGrpc "coinflow/coinflow-server/pkg/grpc"

	"github.com/jinzhu/copier"

	"github.com/google/uuid"
)

// Requests -------------------------------------------

type GetSummaryInPeriodRequestObject struct {
	UserId uuid.UUID
	Begin time.Time
	End time.Time
}

func CreateGetSummaryInPeriodRequestObject(
	ctx context.Context,
	r *pb.GetSummaryInPeriodRequest,
) (*GetSummaryInPeriodRequestObject, error) {
	const op = "CreateGetSummaryInPeriodRequestObject"

	id, err := pkgGrpc.GetUserId(ctx, "user-id")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GetSummaryInPeriodRequestObject{
		UserId: id,
		Begin: r.BeginTime.AsTime(),
		End: r.EndTime.AsTime(),
	}, nil
}

type GetSummaryInLastNMonthsRequestObject struct {
	UserId uuid.UUID
	N int
	CurTime time.Time
	Timezone string
}

func CreateGetSummaryInLastNMonthsRequestObject(
	ctx context.Context,
	r *pb.GetSummaryInLastNMonthsRequest,
) (*GetSummaryInLastNMonthsRequestObject, error) {
	const op = "CreateGetSummaryInLastNMonthsRequestObject"

	id, err := pkgGrpc.GetUserId(ctx, "user-id")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	n := 1

	if r.N < 0 {
		return nil, fmt.Errorf("%s: %w", op, ErrorInvalidMonthsCount)
	} else if r.N != 0 {
		n = int(r.N)
	}

	curTime := time.Now()

	if tm := r.CurTime.AsTime(); !tm.IsZero() {
		curTime = tm
	}

	timezone := "UTC"

	if len(r.Timezone) != 0 {
		timezone = r.Timezone
	}

	return &GetSummaryInLastNMonthsRequestObject{
		UserId: id,
		N: n,
		CurTime: curTime,
		Timezone: timezone,
	}, nil
}

type GetSummaryByCategoriesRequestObject struct {
	UserId uuid.UUID
	Begin time.Time
	End time.Time
}

func CreateGetSummaryByCategoriesRequestObject(
	ctx context.Context,
	r *pb.GetSummaryByCategoriesRequest,
) (*GetSummaryByCategoriesRequestObject, error) {
	const op = "CreateGetSummaryByCategoriesRequestObject"

	id, err := pkgGrpc.GetUserId(ctx, "user-id")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GetSummaryByCategoriesRequestObject{
		UserId: id,
		Begin: r.BeginTime.AsTime(),
		End: r.EndTime.AsTime(),
	}, nil
}

// Responses -------------------------------------------

func GetModelSummaryFromProtobuf(sum *models.Summary) (*pb.Summary, error) {
	const op = "GetModelSummaryFromProtobuf"

	var pbSum pb.Summary

	if err := copier.Copy(&pbSum, sum); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &pbSum, nil
}

func CreateGetSummaryInLastNMonthsResponse(sums []*models.Summary) (*pb.GetSummaryInLastNMonthsResponse, error) {
	const op = "CreateGetSummaryInLastNMonthsResponse"

	var resp pb.GetSummaryInLastNMonthsResponse

	for _, s := range sums {
		pbSum, err := GetModelSummaryFromProtobuf(s)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		resp.Summaries = append(resp.Summaries, pbSum)
	}

	return &resp, nil
}

func CreateGetSummaryByCategoriesResponse(sums map[string]*models.Summary) (*pb.GetSummaryByCategoriesResponse, error) {
	const op = "CreateGetSummaryByCategoriesResponse"

	var resp pb.GetSummaryByCategoriesResponse
	resp.Summaries = map[string]*pb.Summary{}

	for cat, s := range sums {
		pbSum, err := GetModelSummaryFromProtobuf(s)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		resp.Summaries[cat] = pbSum
	}

	return &resp, nil
}
