package types

import (
	pb "coinflow/coinflow-server/gen/storage_service/golang"
	"coinflow/coinflow-server/storage-service/internal/models"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

// Requests -------------------------------------------

const (
	TimeLayout = "02/01/2006 15:04:05 -0700"
)

type GetTransactionsInPeriodRequestObject struct {
	Begin time.Time
	End time.Time
}

func CreateGetTransactionsInPeriodRequestObject(r *pb.GetTransactionsInPeriodRequest) (*GetTransactionsInPeriodRequestObject, error) {
	const op = "CreateGetTransactionsInPeriodRequestObject"

	begin, err := time.Parse(TimeLayout, r.Begin)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	end, err := time.Parse(TimeLayout, r.End)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GetTransactionsInPeriodRequestObject{Begin: begin, End: end}, nil
}

type PostTransactionRequestObject struct {
	Tx *models.Transaction
}

func CreatePostTransactionRequestObject(r *pb.PostTransactionRequest) (*PostTransactionRequestObject, error) {
	const op = "CreatePostTransactionRequestObject"

	var tx models.Transaction

	if err := copier.Copy(&tx, r.Tx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &PostTransactionRequestObject{Tx: &tx}, nil
}

// Responses -------------------------------------------

func CreateGetTransactionsInPeriodResponse(txs []*models.Transaction) (*pb.GetTransactionsInPeriodResponse, error) {
	const op = "CreateGetTransactionsInPeriodResponse"

	var pbTxs []*pb.Transaction

	for _, tx := range txs {
		var pbTx pb.Transaction

		if err := copier.Copy(&pbTx, &tx); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		pbTx.Timestamp = tx.Timestamp.String()

		pbTxs = append(pbTxs, &pbTx)
	}

	return &pb.GetTransactionsInPeriodResponse{Txs: pbTxs}, nil
}

func CreatePostTransactionResponse(txId uuid.UUID) (*pb.PostTransactionResponse, error) {
	return &pb.PostTransactionResponse{TxId: txId.String()}, nil
}
