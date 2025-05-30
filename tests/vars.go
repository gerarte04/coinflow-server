package tests

import (
	tu "coinflow/coinflow-server/pkg/testutils"
	"fmt"
	"net/http"
	"os"
)

const (
	CommitPath = "/commit"
	TransactionPath = "/transaction/id"
	TransactionsInPeriodPath = "/transaction/period"

	RegisterPath = "/auth/register"
	LoginPath = "/auth/login"
	RefreshPath = "/auth/refresh"
	GetUserDataPath = "/user"
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
		"description": "Purchased latte and croissant",
		"category": "food",
		"cost": float64(400),
	}

	exampleInvalidId = "abcdefgh123"

	cookieJar http.CookieJar
	cli *http.Client
)
