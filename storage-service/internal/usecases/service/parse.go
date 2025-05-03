package service

import (
	pb "coinflow/coinflow-server/gen/collection_service/golang"
	"coinflow/coinflow-server/storage-service/internal/models"

	"github.com/jinzhu/copier"
)

func ConvertModelTransactionToProtobuf(tx *models.Transaction) (*pb.Transaction, error) {
	var pbTx pb.Transaction

	if err := copier.Copy(&pbTx, tx); err != nil {
		return nil, err
	}

	return &pbTx, nil
}
