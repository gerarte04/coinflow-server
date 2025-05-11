package middleware

import (
	"coinflow/coinflow-server/pkg/utils"
	pkgCrypto "coinflow/coinflow-server/pkg/utils/crypto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func validate(c *gin.Context, token string, publicKey []byte, place string) error {
	usrId, err := pkgCrypto.ValidateJwtToken(token, publicKey)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": fmt.Sprintf("token from %s: %s", place, err.Error()),
		})

		return err
	}

	c.Header("Subject", usrId.String())
	c.Next()

	return nil
}

func WithAuthMiddleware(publicKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err1 := utils.ParseAuthHeader(c.Request.Header["Authorization"], "Bearer")
		if err1 == nil {
			validate(c, token, publicKey, "header")
			return
		}

		token, err2 := c.Cookie("accessToken")
		if err2 == nil {
			validate(c, token, publicKey, "cookie")
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": fmt.Sprintf("%s: %s; %s", ErrorTokenNotFound, err1, err2),
		})
	}
}
