package types

import (
	"coinflow/coinflow-server/api-gateway/internal/models"
	pb "coinflow/coinflow-server/gen/storage_service/golang"
	"fmt"

	"github.com/jinzhu/copier"
)

// Requests ------------------------------------------------

func CreatePostTransactionRequest(tx *models.Transaction, withAutoCategory bool) (*pb.PostTransactionRequest, error) {
	const op = "CreatePostTransactionRequest"

	var pbTx pb.Transaction

	if err := copier.Copy(&pbTx, tx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &pb.PostTransactionRequest{
		Tx: &pbTx,
		WithAutoCategory: withAutoCategory,
	}, nil
}

// Responses -----------------------------------------------

func CreateGetTransactionResponse(resp *pb.GetTransactionResponse) (*models.Transaction, error) {
	const op = "CreateGetTransactionResponse"

	var tx models.Transaction

	if err := copier.Copy(&tx, resp.Tx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &tx, nil
}

func CreateGetTransactionsInPeriodResponse(resp *pb.GetTransactionsInPeriodResponse) ([]*models.Transaction, error) {
	const op = "CreateGetTransactionsInPeriodResponse"

	txs := make([]*models.Transaction, len(resp.Txs))

	for i, pbTx := range resp.Txs {
		var tx models.Transaction

		if err := copier.Copy(&tx, pbTx); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		txs[i] = &tx
	}

	return txs, nil
}
