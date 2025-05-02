package types

import (
	"coinflow/coinflow-server/collection-service/internal/models"
	pb "coinflow/coinflow-server/gen/collection_service/golang"

	"github.com/jinzhu/copier"
)

// Requests -------------------------------------------

type GetTransactionCategoryRequestObject struct {
	Tx *models.Transaction
}

func CreateGetTransactionCategoryRequestObject(r *pb.GetTransactionCategoryRequest) (*GetTransactionCategoryRequestObject, error) {
	var tx models.Transaction
	
	if err := copier.Copy(&tx, r.Tx); err != nil {
		return nil, err
	}

	return &GetTransactionCategoryRequestObject{Tx: &tx}, nil
}

// Responses -------------------------------------------

func CreateGetTransactionCategoryResponse(category string) (*pb.GetTransactionCategoryResponse, error) {
	return &pb.GetTransactionCategoryResponse{Category: category}, nil
}
