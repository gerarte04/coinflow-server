package http

import "github.com/gin-gonic/gin"

const (
    GetTransactionPath = "/transaction/:ts_id"
    PostTransactionPath = "/commit"
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
