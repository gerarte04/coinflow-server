package types

import (
	"coinflow/coinflow-server/api-gateway/internal/models"
	"coinflow/coinflow-server/pkg/utils"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Transactions ---------------------------------------------

type GetTransactionRequestObject struct {
	TxId 			uuid.UUID
}

func CreateGetTransactionRequestObject(c *gin.Context) (*GetTransactionRequestObject, error) {
	const op = "CreateGetTransactionRequestObject"

	txId, err := utils.ParseStringToUuid(c.Param("tx_id"))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GetTransactionRequestObject{TxId: txId}, nil
}

type GetTransactionsInPeriodRequestObject struct {
	Begin 			string		`json:"begin"`
	End 			string		`json:"end"`
	WithSummary		bool		`json:"with_summary"`
}

func CreateGetTransactionsInPeriodRequestObject(c *gin.Context) (*GetTransactionsInPeriodRequestObject, error) {
	const op = "CreateGetTransactionsInPeriodRequestObject"

	var req GetTransactionsInPeriodRequestObject

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &req, nil
}

type PostTransactionRequestObject struct {
	Tx 	*models.Transaction
}

func CreatePostTransactionRequestObject(c *gin.Context) (*PostTransactionRequestObject, error) {
	const op = "CreatePostTransactionRequestObject"

	var tx models.Transaction

	if err := c.ShouldBindJSON(&tx); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &PostTransactionRequestObject{Tx: &tx}, nil
}

// Users --------------------------------------------------

type LoginRequestObject struct {
	Login		string	`json:"login"`
	Password	string	`json:"password"`
}

func CreateLoginRequestObject(c *gin.Context) (*LoginRequestObject, error) {
	const op = "CreateLoginRequestObject"

	var req LoginRequestObject

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &req, nil
}

type RefreshRequestObject struct {
	RefreshToken	string	`json:"refresh_token"`
}

func CreateRefreshRequestObject(c *gin.Context) (*RefreshRequestObject, error) {
	const op = "CreateRefreshRequestObject"

	var req RefreshRequestObject

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &req, nil
}

type RegisterRequestObject struct {
	User 	models.User 	`json:"usr"`
}

func CreateRegisterRequestObject(c *gin.Context) (*RegisterRequestObject, error) {
	const op = "CreateRegisterRequestObject"

	var req RegisterRequestObject

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &req, nil
}

type GetUserDataRequestObject struct {
	UserId	uuid.UUID
}

func CreateGetUserDataRequestObject(c *gin.Context) (*GetUserDataRequestObject, error) {
	const op = "CreateGetUserDataRequestObject"

	usrId, err := utils.ParseStringToUuid(c.Param("user_id"))

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &GetUserDataRequestObject{UserId: usrId}, nil
}
