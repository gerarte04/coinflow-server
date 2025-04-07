package http

import (
	"coinflow/coinflow-server/restful-api/internal/api/http/types"
	"coinflow/coinflow-server/restful-api/internal/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CoinflowServer struct {
    tsService usecases.TransactionsService
}

func NewCoinflowServer(tsService usecases.TransactionsService) *CoinflowServer {
    return &CoinflowServer{tsService: tsService}
}

// @Summary GetTransaction
// @Description get transaction by id
// @Tags transactions
// @Accept json
// @Produce json
// @Param ts_id path string true "Transaction ID"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Not found"
// @Failure 500 {object} string "Internal error"
// @Router /transaction/{ts_id} [get]
func (s *CoinflowServer) GetTransactionHandler(c *gin.Context) {
    reqObj, err := types.CreateGetTransactionRequestObject(c)
    if err != nil {
        WriteError(c, err)
        return
    }

    res, err := s.tsService.GetTransaction(reqObj.TsId)
    if err != nil {
        WriteError(c, err)
        return
    }

    c.JSON(http.StatusOK, res)
}

// @Summary PostTransaction
// @Description commit transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param ts_id body models.Transaction true "Transaction"
// @Success 201 {object} string "Created"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal error"
// @Router /commit [post]
func (s *CoinflowServer) PostTransactionHandler(c *gin.Context) {
    reqObj, err := types.CreatePostTransactionRequestObject(c)
    if err != nil {
        WriteError(c, err)
        return
    }

    res, err := s.tsService.PostTransaction(reqObj.Ts)
    if err != nil {
        WriteError(c, err)
        return
    }

    c.JSON(http.StatusCreated, gin.H{
        "ts_id": res,
    })
}
