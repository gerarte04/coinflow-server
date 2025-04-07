package types

import (
	"coinflow/coinflow-server/restful-api/internal/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

type GetTransactionRequestObject struct {
    TsId string
}

func CreateGetTransactionRequestObject(c *gin.Context) (*GetTransactionRequestObject, error) {
    tsId, err := ParseStringToTransactionId(c.Param("ts_id"))

    if err != nil {
        return nil, fmt.Errorf("creating request object: %w", err)
    }

    return &GetTransactionRequestObject{TsId: tsId}, nil
}

type PostTransactionRequestObject struct {
    Ts *models.Transaction
}

func CreatePostTransactionRequestObject(c *gin.Context) (*PostTransactionRequestObject, error) {
    var ts models.Transaction

    if err := c.ShouldBindJSON(&ts); err != nil {
        return nil, fmt.Errorf("creating request object: %w", err)
    }

    return &PostTransactionRequestObject{Ts: &ts}, nil
}
