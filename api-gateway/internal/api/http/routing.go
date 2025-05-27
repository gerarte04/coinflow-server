package http

import (
	"coinflow/coinflow-server/api-gateway/internal/middleware"
	pkgHandlers "coinflow/coinflow-server/pkg/http/handlers"

	"github.com/gin-gonic/gin"
)

const (
	GetTransactionPath = "/transaction/id/:tx_id"
	GetTransactionsInPeriodPath = "/transaction/period"
	PostTransactionPath = "/commit"

	LoginPath = "/auth/login"
	RefreshPath = "/auth/refresh"
	RegisterPath = "/auth/register"
	GetUserDataPath = "/user/:user_id"
)

func (s *CoinflowServer) RouteHandlers(engine *gin.Engine, opts ...pkgHandlers.RouterOption) {
	for _, opt := range opts {
		opt(engine)
	}
}

func (s *CoinflowServer) WithAuthServiceHandlers() pkgHandlers.RouterOption {
	return func(engine *gin.Engine) {
		engine.POST(LoginPath, s.LoginHandler)
		engine.POST(RefreshPath, s.RefreshHandler)
		engine.POST(RegisterPath, s.RegisterHandler)
	}
}

func (s *CoinflowServer) WithSecuredUserHandlers(publicKey []byte) pkgHandlers.RouterOption {
	return func(engine *gin.Engine) {
		authGroup := engine.Group("/")
		authGroup.Use(middleware.WithAuthMiddleware(publicKey))

		authGroup.GET(GetTransactionPath, s.GetTransactionHandler)
		authGroup.POST(GetTransactionsInPeriodPath, s.GetTransactionsInPeriodHandler)
		authGroup.POST(PostTransactionPath, s.PostTransactionHandler)
	
		authGroup.GET(GetUserDataPath, s.GetUserDataHandler)
	}
}
