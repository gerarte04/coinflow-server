package handlers

import (
	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

const (
	SwaggerPath = "/swagger/*path"
	HealthCheckPath = "/health"
)

type RouterOption func(engine *gin.Engine)

func WithLogger() RouterOption {
	return func(engine *gin.Engine) {
		engine.Use(gin.LoggerWithWriter(gin.DefaultWriter, HealthCheckPath))
	}
}

func WithRecovery() RouterOption {
	return func(engine *gin.Engine) {
		engine.Use(gin.Recovery())
	}
}

func WithHealthCheck() RouterOption {
	return func(engine *gin.Engine) {
		engine.GET(HealthCheckPath, healthCheckHandler)
	}
}

func WithSwagger() RouterOption {
	return func(engine *gin.Engine) {
		engine.GET(SwaggerPath, gin.WrapF(httpSwagger.WrapHandler))
	}
}
