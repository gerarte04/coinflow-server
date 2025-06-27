package types

import (
	pb "coinflow/coinflow-server/gen/storage_service/golang"
	pkgGrpc "coinflow/coinflow-server/pkg/grpc"
	"coinflow/coinflow-server/pkg/utils"
	"coinflow/coinflow-server/pkg/vars"
	"coinflow/coinflow-server/storage-service/internal/models"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
)

func getUserId(ctx context.Context) (uuid.UUID, error) {
	val, err := pkgGrpc.GetHeader(ctx, "user-id")
	if err != nil {
		return uuid.Nil, err
	}

	id, err := utils.ParseStringToUuid(val)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

// Requests -------------------------------------------

type GetTransactionRequestObject struct {
	UserId uuid.UUID 
	TxId uuid.UUID
}

func CreateGetTransactionRequestObject(ctx context.Context, r *pb.GetTransactionRequest) (*GetTransactionRequestObject, error) {
	const op = "CreateGetTransactionRequestObject"

	
	txId, err := utils.ParseStringToUuid(r.TxId)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	usrId, err := getUserId(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GetTransactionRequestObject{UserId: usrId, TxId: txId}, nil
}

type GetTransactionsInPeriodRequestObject struct {
	Begin time.Time
	End time.Time
	UserId uuid.UUID
}

func CreateGetTransactionsInPeriodRequestObject(ctx context.Context, r *pb.GetTransactionsInPeriodRequest) (*GetTransactionsInPeriodRequestObject, error) {
	const op = "CreateGetTransactionsInPeriodRequestObject"

	begin, err := time.Parse(vars.TimeLayout, r.Begin)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	end, err := time.Parse(vars.TimeLayout, r.End)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	usrId, err := getUserId(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GetTransactionsInPeriodRequestObject{Begin: begin, End: end, UserId: usrId}, nil
}

type PostTransactionRequestObject struct {
	Tx *models.Transaction
	WithAutoCategory bool
}

func CreatePostTransactionRequestObject(ctx context.Context, r *pb.PostTransactionRequest) (*PostTransactionRequestObject, error) {
	const op = "CreatePostTransactionRequestObject"

	var tx models.Transaction

	if err := copier.Copy(&tx, r.Tx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	usrId, err := getUserId(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	tx.UserId = usrId

	return &PostTransactionRequestObject{
		Tx: &tx,
		WithAutoCategory: r.WithAutoCategory,
	}, nil
}

// Responses -------------------------------------------

func CreateGetTransactionResponse(tx *models.Transaction) (*pb.GetTransactionResponse, error) {
	const op = "CreateGetTransactionResponse"

	var pbTx pb.Transaction

	if err := copier.Copy(&pbTx, tx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	pbTx.Timestamp = tx.Timestamp.Format(vars.TimeLayout)

	return &pb.GetTransactionResponse{Tx: &pbTx}, nil
}

func CreateGetTransactionsInPeriodResponse(txs []*models.Transaction) (*pb.GetTransactionsInPeriodResponse, error) {
	const op = "CreateGetTransactionsInPeriodResponse"

	var pbTxs []*pb.Transaction

	for _, tx := range txs {
		var pbTx pb.Transaction

		if err := copier.Copy(&pbTx, &tx); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		pbTx.Timestamp = tx.Timestamp.Format(vars.TimeLayout)

		pbTxs = append(pbTxs, &pbTx)
	}

	return &pb.GetTransactionsInPeriodResponse{Txs: pbTxs}, nil
}
