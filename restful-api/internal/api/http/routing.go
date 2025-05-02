package http

import (
	pkgHandlers "coinflow/coinflow-server/pkg/http/handlers"

	"github.com/gin-gonic/gin"
)

const (
	GetTransactionPath = "/transaction/:ts_id"
	PostTransactionPath = "/commit"
)

func (s *CoinflowServer) RouteHandlers(engine *gin.Engine, opts ...pkgHandlers.RouterOption) {
	for _, opt := range opts {
		opt(engine)
	}
}

func (s *CoinflowServer) WithStandardUserHandlers() pkgHandlers.RouterOption {
	return func(engine *gin.Engine) {
		engine.GET(GetTransactionPath, s.GetTransactionHandler)
		engine.POST(PostTransactionPath, s.PostTransactionHandler)
	}
}
