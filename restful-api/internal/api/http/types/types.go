package types

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"path"
)

// Requests -------------------------------------------

type GetTransactionRequestObject struct {
    TsId string
}

func CreateGetTransactionRequestObject(r *http.Request) (*GetTransactionRequestObject, error) {
    tsId, err := ParseStringToTransactionId(path.Base(r.URL.Path))

    if err != nil {
        return nil, fmt.Errorf("creating request object: %w", err)
    }

    return &GetTransactionRequestObject{TsId: tsId}, nil
}

type PostTransactionRequestObject struct {
    Ts *models.Transaction
}

func CreatePostTransactionRequestObject(r *http.Request) (*PostTransactionRequestObject, error) {
    data, err := io.ReadAll(r.Body)

    if err != nil {
        return nil, fmt.Errorf("creating request object: %w", err)
    }

    ts, err := ParseByteToTransaction(data)

    if err != nil {
        return nil, fmt.Errorf("creating request object: %w", err)
    }

    return &PostTransactionRequestObject{Ts: ts}, nil
}

// Responses -------------------------------------------

func CreateGetTransactionResponse(ts *models.Transaction) ([]byte, error) {
    data, err := ParseTransactionToByte(ts)

    if err != nil {
        return nil, fmt.Errorf("creating response object: %w", err)
    }

    return data, nil
}

func CreatePostTransactionResponse(ts *models.Transaction) ([]byte, error) {
    mp := map[string]string{
        "ts_id": ts.Id,
    }

    data, err := json.Marshal(mp)

    if err != nil {
        return nil, fmt.Errorf("creating response object: %w", err)
    }

    return data, nil
}
