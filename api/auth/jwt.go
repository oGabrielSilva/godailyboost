package auth

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	Uid         string      `json:"uid"`
	Email       string      `json:"email"`
	Name        string      `json:"name"`
	Authorities []Authority `json:"authorities"`
	jwt.StandardClaims
}

var (
	secret []byte
)

func ConfigureJWT() {
	e := os.Getenv("JWT_SECRET")
	if e == "" {
		panic("env JWT_SECRET not found")
	}

	secret = []byte(e)
}

func RecoveryToken(c *gin.Context) (claims *Claims, ok bool) {
	authorizationHeader := c.GetHeader("Authorization")
	if authorizationHeader == "" {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok = token.Claims.(*Claims)
	return
}
