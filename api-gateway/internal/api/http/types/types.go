package types

import (
	"coinflow/coinflow-server/api-gateway/internal/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetTransactionsInPeriodRequestObject struct {
	Begin 			string		`json:"begin"`
	End 			string		`json:"end"`
	WithSummary		bool		`json:"with_summary"`
}

func CreateGetTransactionsInPeriodRequestObject(c *gin.Context) (*GetTransactionsInPeriodRequestObject, error) {
	const op = "CreateGetTransactionRequestObject"

	var req GetTransactionsInPeriodRequestObject

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &req, nil
}

type PostTransactionRequestObject struct {
	Tx *models.Transaction
}

func CreatePostTransactionRequestObject(c *gin.Context) (*PostTransactionRequestObject, error) {
	const op = "CreatePostTransactionRequestObject"

	var tx models.Transaction

	if err := c.ShouldBindJSON(&tx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &PostTransactionRequestObject{Tx: &tx}, nil
}
