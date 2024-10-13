package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/oGabrielSilva/godailyboost/api/routes"
)

type SessionDTO struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func SignUp() gin.HandlerFunc {
	return routes.ExceptionHandler(func(c *gin.Context) error {
		return nil
	})
}
