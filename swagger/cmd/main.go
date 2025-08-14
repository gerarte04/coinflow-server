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

// @Summary Get transaction by id
// @Description WARNING: Viewing other user's transactions is not allowed.
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

// @Summary Get transactions within time interval
// @Description Time should be presented in RFC3339 format.
// @Tags transactions
// @Accept json
// @Produce json
// @Param user_id query string true "User ID"
// @Param begin_time query string true "Begin time in RFC3339 format"
// @Param end_time query string true "End time in RFC3339 format"
// @Param page_size query int false "Requested page size (optional, default is 10)"
// @Param page_token query string false "Last timestamp used in previous page (optional, for keyset pagination)"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 404 {object} string "Not found"
// @Failure 500 {object} string "Internal error"
// @Router /transaction/period [get]
func ListTransactions() {}

// @Summary Commit transaction
// @Description Transaction's type and category must be one of allowed values (view docs/VARS.md on github for values list).
// @Tags transactions
// @Accept json
// @Produce json
// @Param tx body models.Transaction true "Transaction data"
// @Param user_id query string true "Creator id"
// @Param with_auto_category query bool false "Defines if category should be auto-detected by description (default is false)"
// @Success 201 {object} string "Created"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal error"
// @Router /commit [post]
func CreateTransaction() {}

// Summary -------------------------------------------

// @Summary Get summary of transactions within time interval
// @Tags summary
// @Produce json
// @Param user_id query string true "User ID"
// @Param begin_time query string true "Begin time in RFC3339"
// @Param end_time query string true "End time in RFC3339"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 500 {object} string "Internal error"
// @Router /summary/period [get]
func GetSummaryInPeriod() {}

// @Summary Get summary of transactions for last N months
// @Tags summary
// @Produce json
// @Param user_id query string true "User ID"
// @Param n query string false "Count of last months (non-negative, optional, default is 1)"
// @Param cur_time query string false "Current in RFC3339, optional"
// @Param timezone query string false "Supported timezone (ex. Europe/Moscow, optional, default is UTC)"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 500 {object} string "Internal error"
// @Router /summary/last-months [get]
func GetSummaryInLastNMonths() {}

// @Summary Get summary of transactions for different categories (within time interval)
// @Tags summary
// @Produce json
// @Param user_id query string true "User ID"
// @Param begin_time query string true "Begin time in RFC3339"
// @Param end_time query string true "End time in RFC3339"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 403 {object} string "Forbidden"
// @Failure 500 {object} string "Internal error"
// @Router /summary/categories [get]
func GetSummaryByCategories() {}

// Users ---------------------------------------------

// @Summary Login and get pair of tokens
// @Tags users
// @Accept json
// @Produce json
// @Param reqObj body types.LoginRequestObject true "Login and password"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal error"
// @Router /auth/login [post]
func Login() {}

// @Summary Refresh and get new pair of tokens
// @Tags users
// @Accept json
// @Produce json
// @Param reqObj body types.RefreshRequestObject true "Refresh token"
// @Success 200 {object} string "OK"
// @Failure 400 {object} string "Bad request"
// @Failure 401 {object} string "Unauthorized"
// @Failure 500 {object} string "Internal error"
// @Router /auth/refresh [post]
func Refresh() {}

// @Summary Register new user
// @Tags users
// @Accept json
// @Produce json
// @Param reqObj body models.User true "User data"
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
// @version 0.1.0

// @host localhost:8080
// @BasePath /v1
func main() {
	var cfg config.SwaggerConfig
	flg := pkgConfig.ParseFlags()
	pkgConfig.MustLoadConfig(flg.ConfigPath, &cfg)

	r := chi.NewRouter()
	r.Get(cfg.SwaggerPath+"/*", httpSwagger.WrapHandler)

	addr := fmt.Sprintf("%s:%s", cfg.HttpCfg.Host, cfg.HttpCfg.Port)

	if err := server.CreateServer(addr, r); err != nil {
		log.Fatalf("failed to run server: %s", err.Error())
	}
}
