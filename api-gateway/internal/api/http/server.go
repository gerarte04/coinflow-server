package http

import (
	"coinflow/coinflow-server/api-gateway/config"
	"coinflow/coinflow-server/api-gateway/internal/api/http/types"
	"coinflow/coinflow-server/api-gateway/internal/clients/grpc"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CoinflowServer struct {
	storageCli grpc.StorageClient
	collectCli grpc.CollectionClient
	authCli grpc.AuthClient
	securityCfg config.SecurityConfig
}

func NewCoinflowServer(
	storageCli grpc.StorageClient,
	collectCli grpc.CollectionClient,
	authCli grpc.AuthClient,
	securityCfg config.SecurityConfig,
) *CoinflowServer {
	return &CoinflowServer{
		storageCli: storageCli,
		collectCli: collectCli,
		authCli: authCli,
		securityCfg: securityCfg,
	}
}

// Transactions ----------------------------------------------

// @Summary GetTransaction
// @Description get transaction by id
// @Tags transactions
// @Accept json
// @Produce json
// @Param tx_id path string true "Transaction ID"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 404 {object} string "Not found"
// @Failure 500 {object} string "Internal error"
// @Router /transaction/id/{tx_id} [get]
func (s *CoinflowServer) GetTransactionHandler(c *gin.Context) {
	reqObj, err := types.CreateGetTransactionRequestObject(c)
	if err != nil {
		WriteParseError(c, err)
		return
	}

	res, err := s.storageCli.GetTransaction(c.Request.Context(), reqObj.UserId, reqObj.TxId)
	if err != nil {
		WriteGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

// @Summary GetTransactionsInPeriod
// @Description get transactions in period between begin and end
// @Tags transactions
// @Accept json
// @Produce json
// @Param reqObj body types.GetTransactionsInPeriodRequestObject true "Request object"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 404 {object} string "Not found"
// @Failure 500 {object} string "Internal error"
// @Router /transaction/period [post]
func (s *CoinflowServer) GetTransactionsInPeriodHandler(c *gin.Context) {
	reqObj, err := types.CreateGetTransactionsInPeriodRequestObject(c)
	if err != nil {
		WriteParseError(c, err)
		return
	}

	res, err := s.storageCli.GetTransactionsInPeriod(c.Request.Context(), reqObj.UserId, reqObj.Begin, reqObj.End)
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
// @Param tx body types.PostTransactionRequestObject true "transaction"
// @Success 201 {object} string "Created"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal error"
// @Router /commit [post]
func (s *CoinflowServer) PostTransactionHandler(c *gin.Context) {
	reqObj, err := types.CreatePostTransactionRequestObject(c)
	if err != nil {
		WriteParseError(c, err)
		return
	}

	res, err := s.storageCli.PostTransaction(c.Request.Context(), reqObj.Tx, reqObj.WithAutoCategory)
	if err != nil {
		WriteGrpcError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tx_id": res,
	})
}

// Users ---------------------------------------------

func (s *CoinflowServer) setAccessToken(c *gin.Context, token string) {
	c.SetCookie("accessToken", token,
		int(math.Ceil(s.securityCfg.AccessExpirationTime.Seconds())),
		"",
		"",
		!s.securityCfg.AllowUnsecureCookies,
		true,
	)
}

// @Summary Login
// @Description login and get tokens
// @Tags users
// @Accept json
// @Produce json
// @Param reqObj body types.LoginRequestObject true "request object"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal error"
// @Router /auth/login [post]
func (s *CoinflowServer) LoginHandler(c *gin.Context) {
	reqObj, err := types.CreateLoginRequestObject(c)
	if err != nil {
		WriteParseError(c, err)
		return
	}

	access, refresh, err := s.authCli.Login(c.Request.Context(), reqObj.Login, reqObj.Password)
	if err != nil {
		WriteGrpcError(c, err)
		return
	}

	s.setAccessToken(c, access)

	c.JSON(http.StatusOK, gin.H{
		"access_token": access,
		"refresh_token": refresh,
	})
}

// @Summary Refresh
// @Description refresh and get new tokens
// @Tags users
// @Accept json
// @Produce json
// @Param reqObj body types.RefreshRequestObject true "request object"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal error"
// @Router /auth/refresh [post]
func (s *CoinflowServer) RefreshHandler(c *gin.Context) {
	reqObj, err := types.CreateRefreshRequestObject(c)
	if err != nil {
		WriteParseError(c, err)
		return
	}

	access, refresh, err := s.authCli.Refresh(c.Request.Context(), reqObj.RefreshToken)
	if err != nil {
		WriteGrpcError(c, err)
		return
	}

	s.setAccessToken(c, access)

	c.JSON(http.StatusOK, gin.H{
		"access_token": access,
		"refresh_token": refresh,
	})
}

// @Summary Register
// @Description register new user
// @Tags users
// @Accept json
// @Produce json
// @Param reqObj body types.RegisterRequestObject true "request object"
// @Success 201 {object} string "Created"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal error"
// @Router /auth/register [post]
func (s *CoinflowServer) RegisterHandler(c *gin.Context) {
	reqObj, err := types.CreateRegisterRequestObject(c)
	if err != nil {
		WriteParseError(c, err)
		return
	}

	res, err := s.authCli.Register(c.Request.Context(), &reqObj.User)
	if err != nil {
		WriteGrpcError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id": res,
	})
}

// @Summary GetUserData
// @Description get user data by id
// @Tags users
// @Accept json
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal error"
// @Router /user/{user_id} [get]
func (s *CoinflowServer) GetUserDataHandler(c *gin.Context) {
	reqObj, err := types.CreateGetUserDataRequestObject(c)
	if err != nil {
		WriteParseError(c, err)
		return
	}

	res, err := s.authCli.GetUserData(c.Request.Context(), reqObj.UserId)
	if err != nil {
		WriteGrpcError(c, err)
		return
	}

	c.JSON(http.StatusOK, res)
}
