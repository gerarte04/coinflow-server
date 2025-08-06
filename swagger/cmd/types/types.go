package types

import (
	"coinflow/coinflow-server/swagger/cmd/types/models"

	"github.com/google/uuid"
)

type RefreshRequestObject struct {
	RefreshToken string `json:"refresh_token"`
}

type LoginRequestObject struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type GetUserDataRequestObject struct {
	UserId uuid.UUID `json:"user_id"`
}

type GetTransactionRequestObject struct {
	TxId uuid.UUID `json:"tx_id"`
}

type GetTransactionsInPeriodRequestObject struct {
	Begin       string `json:"begin"`
	End         string `json:"end"`
	WithSummary bool   `json:"with_summary"`
}

type PostTransactionRequestObject struct {
	Tx               *models.Transaction `json:"tx"`
	WithAutoCategory bool                `json:"with_auto_category"`
}
