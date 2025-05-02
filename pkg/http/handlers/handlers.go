package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func healthCheckHandler(c *gin.Context) {
	c.String(http.StatusOK, "successful healthcheck")
}
