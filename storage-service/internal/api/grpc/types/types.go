package types

import (
	pb "coinflow/coinflow-server/gen/storage_service/golang"
	pkgGrpc "coinflow/coinflow-server/pkg/grpc"
	"coinflow/coinflow-server/pkg/utils"
	"coinflow/coinflow-server/storage-service/config"
	"coinflow/coinflow-server/storage-service/internal/models"
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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

	usrId, err := pkgGrpc.GetUserId(ctx, "user-id")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GetTransactionRequestObject{UserId: usrId, TxId: txId}, nil
}

type ListTransactionsRequestObject struct {
	Begin time.Time
	End time.Time
	UserId uuid.UUID

	PageSize int
}

func CreateListTransactionsRequestObject(
	ctx context.Context,
	r *pb.ListTransactionsRequest,
	cfg config.ServiceConfig,
) (*ListTransactionsRequestObject, error) {
	const op = "CreateListTransactionsRequestObject"

	usrId, err := pkgGrpc.GetUserId(ctx, "user-id")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	pageSize := cfg.DefaultPageSize

	if r.PageSize < 0 {
		return nil, fmt.Errorf("%s: %w", op, ErrorInvalidPageSize)
	} else if r.PageSize != 0 {
		pageSize = int(r.PageSize)
	}

	begin, end := r.BeginTime.AsTime(), r.EndTime.AsTime()

	pageToken, err := url.QueryUnescape(r.PageToken)
	if err != nil {
		return nil, fmt.Errorf("%s: %s: %w", op, ErrorInvalidPageToken.Error(), err)
	}

	if len(pageToken) > 0 {
		if lt, err := time.Parse(time.RFC3339, pageToken); err != nil {
			return nil, fmt.Errorf("%s: %s: %w", op, ErrorInvalidPageToken.Error(), err)
		} else {
			end = lt
		}
	}

	return &ListTransactionsRequestObject{
		Begin: begin,
		End: end,
		UserId: usrId,
		PageSize: pageSize,
	}, nil
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

	usrId, err := pkgGrpc.GetUserId(ctx, "user-id")
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

	pbTx.Timestamp = timestamppb.New(tx.Timestamp)

	return &pbTx, nil
}

func CreateListTransactionsResponse(txs []*models.Transaction) (*pb.ListTransactionsResponse, error) {
	const op = "CreateListTransactionsResponse"

	if len(txs) == 0 {
		return &pb.ListTransactionsResponse{}, nil
	}

	var pbTxs []*pb.Transaction
	nextToken := txs[0].Timestamp

	for _, tx := range txs {
		pbTx, err := GetProtobufTxFromModel(tx)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		pbTxs = append(pbTxs, pbTx)

		if nextToken.After(tx.Timestamp) {
			nextToken = tx.Timestamp
		}
	}

	return &pb.ListTransactionsResponse{
		Txs: pbTxs,
		NextPageToken: nextToken.Format(time.RFC3339Nano),
	}, nil
}
