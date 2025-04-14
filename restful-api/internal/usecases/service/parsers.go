package service

import (
	pb "coinflow/coinflow-server/gen/collect_service"
	"coinflow/coinflow-server/restful-api/internal/models"

	"github.com/jinzhu/copier"
)

func ConvertModelTransactionToProtobuf(ts *models.Transaction) (*pb.Transaction, error) {
	var pbTs pb.Transaction

	if err := copier.Copy(&pbTs, ts); err != nil {
		return nil, err
	}

	return &pbTs, nil
}
