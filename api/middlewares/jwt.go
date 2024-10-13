package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oGabrielSilva/godailyboost/api/auth"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := auth.RecoveryToken(c)
		if !ok {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
