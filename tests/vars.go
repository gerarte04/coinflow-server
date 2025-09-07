package tests

import (
	tu "coinflow/coinflow-server/pkg/testutils"
	"fmt"
	"net/http"
	"os"
)

const (
	CommitPath = "/v1/commit"
	TransactionPath = "/v1/transaction/id"
	TransactionsInPeriodPath = "/v1/transaction/period"

	RegisterPath = "/v1/auth/register"
	LoginPath = "/v1/auth/login"
	RefreshPath = "/v1/auth/refresh"
	GetUserDataPath = "/v1/user"
)

var (
	addr = fmt.Sprintf("http://%s:%s", os.Getenv("HTTP_HOST"), os.Getenv("HTTP_PORT"))

	exampleUser = tu.Payload{
		"email": "johndoe@gmail.com",
		"login": "johndoe",
		"name": "John Doe",
		"password": "pass",
		"phone": "+12345678900",
	}

	exampleTx = tu.Payload{
		"type": "purchase",
		"target": "Coffee Point",
		"description": "Latte and croissant",
		"category": "food",
		"cost": float64(400),
	}

	exampleInvalidId = "abcdefgh123"

	cookieJar http.CookieJar
	cli *http.Client
)
