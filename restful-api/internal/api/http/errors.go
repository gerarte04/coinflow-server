package http

import (
	"coinflow/coinflow-server/restful-api/internal/api/http/types"
	"coinflow/coinflow-server/restful-api/internal/repository"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	errorCodes = map[error]int{
		types.ErrorInvalidId: http.StatusBadRequest,
		types.ErrorParseTransaction: http.StatusBadRequest,

		repository.ErrorTransactionKeyNotFound: http.StatusNotFound,
		repository.ErrorUserKeyNotFound: http.StatusNotFound,
		repository.ErrorNoSuchCredExists: http.StatusForbidden,
	}

	codeMessages = map[int]string{
		http.StatusBadRequest: "bad request",
		http.StatusInternalServerError: "internal error",
		http.StatusForbidden: "forbidden",
		http.StatusNotFound: "not found",
	}
)

func WriteError(c *gin.Context, err error) {
	log.Printf("%s\n", err.Error())

	var basicErr error
	for next := err; next != nil; next = errors.Unwrap(basicErr) {
		basicErr = next
	}

	errCode, ok := errorCodes[basicErr]
	var message string

	if !ok {
		log.Printf("warning: undocumented wrapped error\n")
		errCode = http.StatusInternalServerError
	}
	
	if message, ok = codeMessages[errCode]; !ok {
		log.Printf("warning: undocumented status code\n")
		message = "\\undocumented status\\"
	}

	c.JSON(errCode, gin.H{
		"error": fmt.Sprintf("%s: %s", message, err.Error()),
	})
}
