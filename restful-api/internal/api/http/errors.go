package http

import (
	"coinflow/coinflow-server/restful-api/internal/api/http/types"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var (
    errorCodes = map[error]int{
        types.ErrorInvalidId: http.StatusBadRequest,
        types.ErrorParseTransaction: http.StatusBadRequest,
        nil: http.StatusInternalServerError,
    }

    codeMessages = map[int]string{
        http.StatusBadRequest: "bad request",
        http.StatusInternalServerError: "internal error",
        http.StatusForbidden: "forbidden",
    }
)

func WriteError(w http.ResponseWriter, err error) {
    basicErr := errors.Unwrap(err)
    log.Printf("%s\n", err.Error())

    if _, ok := errorCodes[basicErr]; !ok {
        log.Printf("failed to write error: undocumented wrapped error\n")
        return
    } else if _, ok := codeMessages[errorCodes[basicErr]]; !ok {
        log.Printf("failed to write error: undocumented status code\n")
        return
    }

    errCode := errorCodes[basicErr]
    http.Error(w, fmt.Sprintf("%s: %s", codeMessages[errCode], err.Error()), errCode)
}
