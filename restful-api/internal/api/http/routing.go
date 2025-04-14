package http

import (
	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	GetTransactionPath = "/transaction/:ts_id"
	PostTransactionPath = "/commit"

	SwaggerPath = "/swagger/*path"
)

type RouterOption func(engine *gin.Engine)

func (s *CoinflowServer) RouteHandlers(engine *gin.Engine, opts ...RouterOption) {
	for _, opt := range opts {
		opt(engine)
	}
}

func (s *CoinflowServer) WithStandardUserHandlers() RouterOption {
	return func(engine *gin.Engine) {
		engine.GET(GetTransactionPath, s.GetTransactionHandler)
		engine.POST(PostTransactionPath, s.PostTransactionHandler)
	}
}

func (s *CoinflowServer) WithSwagger() RouterOption {
	return func(engine *gin.Engine) {
		engine.GET(SwaggerPath, gin.WrapF(httpSwagger.WrapHandler))
	}
}
