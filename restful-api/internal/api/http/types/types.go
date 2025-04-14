package types

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GetTransactionRequestObject struct {
	TsId uuid.UUID
}

func CreateGetTransactionRequestObject(c *gin.Context) (*GetTransactionRequestObject, error) {
	const fc = "CreateGetTransactionRequestObject"

	tsId, err := ParseStringToTransactionId(c.Param("ts_id"))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", fc, err)
	}

	return &GetTransactionRequestObject{TsId: tsId}, nil
}

type PostTransactionRequestObject struct {
	Ts *models.Transaction
}

func CreatePostTransactionRequestObject(c *gin.Context) (*PostTransactionRequestObject, error) {
	const fc = "CreateGetTransactionRequestObject"

	var ts models.Transaction

	if err := c.ShouldBindJSON(&ts); err != nil {
		return nil, fmt.Errorf("%s: %w", fc, err)
	}

	return &PostTransactionRequestObject{Ts: &ts}, nil
}
