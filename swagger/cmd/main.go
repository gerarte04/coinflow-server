package main

import (
	pkgConfig "coinflow/coinflow-server/pkg/config"
	"coinflow/coinflow-server/pkg/http/server"
	"coinflow/coinflow-server/swagger/config"
	"fmt"
	"log"

	_ "coinflow/coinflow-server/swagger/docs"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Transactions ----------------------------------------------

// @Summary GetTransaction
// @Description get transaction by id
// @Tags transactions
// @Produce json
// @Param tx_id path string true "Transaction ID"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 404 {object} string "Not found"
// @Failure 500 {object} string "Internal error"
// @Router /transaction/id/{tx_id} [get]
func GetTransaction() {}

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
func GetTransactionsInPeriod() {}

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
func PostTransaction() {}

// Users ---------------------------------------------

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
func Login() {}

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
func Refresh() {}

// @Summary Register
// @Description register new user
// @Tags users
// @Accept json
// @Produce json
// @Param reqObj body models.User true "request object"
// @Success 201 {object} string "Created"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal error"
// @Router /auth/register [post]
func Register() {}

// @Summary GetUserData
// @Description get user data by id
// @Tags users
// @Produce json
// @Param user_id path string true "User ID"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal error"
// @Router /user/{user_id} [get]
func GetUserData() {}

// @title Coinflow API
// @version 1.0
// @description API Gateway for Coinflow service

// @host localhost:8080
// @BasePath /v1
func main() {
	var cfg config.SwaggerConfig
	flg := pkgConfig.ParseFlags()
	pkgConfig.MustLoadConfig(flg.ConfigPath, &cfg)

	r := chi.NewRouter()
	r.Get(cfg.SwaggerPath + "/*", httpSwagger.WrapHandler)

	addr := fmt.Sprintf("%s:%s", cfg.HttpCfg.Host, cfg.HttpCfg.Port)

	if err := server.CreateServer(addr, r); err != nil {
		log.Fatalf("failed to run server: %s", err.Error())
	}
}
