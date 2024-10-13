package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oGabrielSilva/godailyboost/api/db"
	"github.com/oGabrielSilva/godailyboost/api/middlewares"
)

func MapApiRoutesV1(r *gin.Engine, db *db.Handler) {
	v1 := r.Group("/v1")

	v1.GET("/health", ExceptionHandler(func(ctx *gin.Context) error {
		ctx.JSON(http.StatusOK, gin.H{"health": "OK"})
		return nil
	}))

	v1.POST("/session")

	authV1 := v1.Group("/")

	authV1.Use(middlewares.JWTMiddleware())
}
