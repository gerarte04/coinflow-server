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

type ListTransactionsRequestObject struct {
	Begin time.Time
	End time.Time
	UserId uuid.UUID
}

func CreateListTransactionsRequestObject(ctx context.Context, r *pb.ListTransactionsRequest) (*ListTransactionsRequestObject, error) {
	const op = "CreateGetTransactionsInPeriodRequestObject"

	begin, err := time.Parse(time.RFC3339, r.BeginTime)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	end, err := time.Parse(time.RFC3339, r.EndTime)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	usrId, err := getUserId(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &ListTransactionsRequestObject{Begin: begin, End: end, UserId: usrId}, nil
}

type PostTransactionRequestObject struct {
	Tx *models.Transaction
	WithAutoCategory bool
}

func MakeCreateTransactionRequestObject(ctx context.Context, r *pb.CreateTransactionRequest) (*PostTransactionRequestObject, error) {
	const op = "MakeCreateTransactionRequestObject"

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

func GetProtobufTxFromModel(tx *models.Transaction) (*pb.Transaction, error) {
	const op = "GetProtobufTxFromModel"

	var pbTx pb.Transaction

	if err := copier.Copy(&pbTx, tx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	pbTx.Timestamp = tx.Timestamp.Format(vars.TimeLayout)

	return &pbTx, nil
}

func CreateListTransactionsResponse(txs []*models.Transaction) (*pb.ListTransactionsResponse, error) {
	const op = "CreateListTransactionsResponse"

	var pbTxs []*pb.Transaction

	for _, tx := range txs {
		pbTx, err := GetProtobufTxFromModel(tx)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		pbTxs = append(pbTxs, pbTx)
	}

	return &pb.ListTransactionsResponse{Txs: pbTxs}, nil
}
