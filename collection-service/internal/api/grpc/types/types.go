package types

import (
	"coinflow/coinflow-server/collection-service/internal/models"
	pb "coinflow/coinflow-server/gen/collection_service/golang"

	"github.com/jinzhu/copier"
)

// Requests -------------------------------------------

type GetTransactionCategoryRequestObject struct {
	Ts *models.Transaction
}

func CreateGetTransactionCategoryRequestObject(r *pb.GetTransactionCategoryRequest) (*GetTransactionCategoryRequestObject, error) {
	var ts models.Transaction
	
	if err := copier.Copy(&ts, r.Ts); err != nil {
		return nil, err
	}

	return &GetTransactionCategoryRequestObject{Ts: &ts}, nil
}

// Responses -------------------------------------------

func CreateGetTransactionCategoryResponse(category string) (*pb.GetTransactionCategoryResponse, error) {
	return &pb.GetTransactionCategoryResponse{Category: category}, nil
}
