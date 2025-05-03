package http

import (
	"coinflow/coinflow-server/api-gateway/internal/api/http/types"
	"coinflow/coinflow-server/api-gateway/internal/clients/grpc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CoinflowServer struct {
	storageCli grpc.StorageClient
	collectCli grpc.CollectionClient
}

func NewCoinflowServer(storageCli grpc.StorageClient, collectCli grpc.CollectionClient) *CoinflowServer {
	return &CoinflowServer{
		storageCli: storageCli,
		collectCli: collectCli,
	}
}

// @Summary GetTransactionsInPeriod
// @Description get transactions in period between begin and end
// @Tags transactions
// @Accept json
// @Produce json
// @Param reqObj body types.GetTransactionsInPeriodRequestObject true "Request object"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 404 {object} string "Not found"
// @Failure 500 {object} string "Internal error"
// @Router /transaction/period [post]
func (s *CoinflowServer) GetTransactionsInPeriodHandler(c *gin.Context) {
	reqObj, err := types.CreateGetTransactionsInPeriodRequestObject(c)
	if err != nil {
		WriteParseError(c, err)
		return
	}

	res, err := s.storageCli.GetTransactionsInPeriod(reqObj.Begin, reqObj.End)
	if err != nil {
		WriteGrpcError(c, err)
		return
	}

	body := gin.H{
		"txs": res,
	}

	if reqObj.WithSummary {
		body["summary"] = "contains"
	}

	c.JSON(http.StatusOK, body)
}

// @Summary PostTransaction
// @Description commit transaction
// @Tags transactions
// @Accept json
// @Produce json
// @Param tx body models.Transaction true "Transaction"
// @Success 201 {object} string "Created"
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal error"
// @Router /commit [post]
func (s *CoinflowServer) PostTransactionHandler(c *gin.Context) {
	reqObj, err := types.CreatePostTransactionRequestObject(c)
	if err != nil {
		WriteParseError(c, err)
		return
	}

	res, err := s.storageCli.PostTransaction(reqObj.Tx)
	if err != nil {
		WriteGrpcError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tx_id": res,
	})
}
