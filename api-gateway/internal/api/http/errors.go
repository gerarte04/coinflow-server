package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	statusCodes = map[codes.Code]int{
		codes.InvalidArgument: http.StatusBadRequest,
		codes.NotFound: http.StatusNotFound,
		codes.PermissionDenied: http.StatusForbidden,
		codes.Internal: http.StatusInternalServerError,
		codes.Unauthenticated: http.StatusUnauthorized,
	}

	codeMessages = map[int]string{
		http.StatusBadRequest: "bad request",
		http.StatusInternalServerError: "internal error",
		http.StatusForbidden: "forbidden",
		http.StatusNotFound: "not found",
		http.StatusUnauthorized: "unauthorized",
	}
)

func WriteGrpcError(c *gin.Context, err error) {
	grpcCode := status.Code(err)
	httpCode, ok := statusCodes[grpcCode]

	if !ok {
		httpCode = http.StatusInternalServerError
	}

	message, ok := codeMessages[httpCode]
	
	if !ok {
		message = "\\undocumented status\\"
	}

	c.AbortWithStatusJSON(httpCode, gin.H{
		"error": fmt.Sprintf("%s: %s", message, err.Error()),
	})
}

func WriteParseError(c *gin.Context, err error) {
	log.Printf("%s", err.Error())

	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": fmt.Sprintf("internal error: %s", err.Error()),
	})
}
