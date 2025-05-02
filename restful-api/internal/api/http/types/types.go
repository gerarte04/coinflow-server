package types

import (
	"coinflow/coinflow-server/pkg/utils"
	"coinflow/coinflow-server/restful-api/internal/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetTransactionRequestObject struct {
	TxId uuid.UUID
}

func CreateGetTransactionRequestObject(c *gin.Context) (*GetTransactionRequestObject, error) {
	const fc = "CreateGetTransactionRequestObject"

	txId, err := utils.ParseStringToTransactionId(c.Param("ts_id"))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", fc, err)
	}

	return &GetTransactionRequestObject{TxId: txId}, nil
}

type PostTransactionRequestObject struct {
	Tx *models.Transaction
}

func CreatePostTransactionRequestObject(c *gin.Context) (*PostTransactionRequestObject, error) {
	const fc = "CreateGetTransactionRequestObject"

	var tx models.Transaction

	if err := c.ShouldBindJSON(&tx); err != nil {
		return nil, fmt.Errorf("%s: %w", fc, err)
	}

	return &PostTransactionRequestObject{Tx: &tx}, nil
}
