package http

import (
	pkgHandlers "coinflow/coinflow-server/pkg/http/handlers"

	"github.com/gin-gonic/gin"
)

const (
	GetTransactionsInPeriodPath = "/transaction/period"
	PostTransactionPath = "/commit"
)

func (s *CoinflowServer) RouteHandlers(engine *gin.Engine, opts ...pkgHandlers.RouterOption) {
	for _, opt := range opts {
		opt(engine)
	}
}

func (s *CoinflowServer) WithStandardUserHandlers() pkgHandlers.RouterOption {
	return func(engine *gin.Engine) {
		engine.POST(GetTransactionsInPeriodPath, s.GetTransactionsInPeriodHandler)
		engine.POST(PostTransactionPath, s.PostTransactionHandler)
	}
}
