package types

import (
	pb "coinflow/coinflow-server/gen/cfapi"
	"coinflow/coinflow-server/internal/models"

	"github.com/jinzhu/copier"
)

// Requests -------------------------------------------

type GetTransactionRequestObject struct {
    TsId string
}

func CreateGetTransactionRequestObject(r *pb.GetTransactionRequest) (*GetTransactionRequestObject, error) {
    tsId := r.GetTsId()

    if tsId == "" {
        return nil, ErrorEmptyId
    }

    return &GetTransactionRequestObject{TsId: tsId}, nil
}

type PostTransactionRequestObject struct {
    Ts *models.Transaction
}

func CreatePostTransactionRequestObject(r *pb.PostTransactionRequest) (*PostTransactionRequestObject, error) {
    var ts models.Transaction
    
    if err := copier.Copy(&ts, r.Ts); err != nil {
        return nil, err
    }

    return &PostTransactionRequestObject{Ts: &ts}, nil
}

// Responses -------------------------------------------

func CreateGetTransactionResponse(ts *models.Transaction) (*pb.GetTransactionResponse, error) {
    var pbTs pb.Transaction

    if err := copier.Copy(&pbTs, ts); err != nil {
        return nil, err
    }

    return &pb.GetTransactionResponse{Ts: &pbTs}, nil
}
