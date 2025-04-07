package types

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

func ParseStringToTransactionId(s string) (string, error) {
    if _, err := uuid.Parse(s); err != nil {
        return "", fmt.Errorf("%w: %s", ErrorInvalidId, err.Error())
    }

    return s, nil
}

func ParseByteToTransaction(data []byte) (*models.Transaction, error) {
    var ts models.Transaction

    if err := json.Unmarshal(data, &ts); err != nil {
        return nil, fmt.Errorf("%w: %s", ErrorParseTransaction, err.Error())
    }

    return &ts, nil
}

func ParseTransactionToByte(ts *models.Transaction) ([]byte, error) {
    data, err := json.Marshal(ts)
    
    if err != nil {
        return nil, fmt.Errorf("%w: %s", ErrorParseTransaction, err.Error())
    }

    return data, nil
}
