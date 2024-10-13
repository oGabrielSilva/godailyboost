package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/oGabrielSilva/godailyboost/api/exceptions"
)

func ExceptionHandler(handler func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := handler(c); err != nil {
			if ex, ok := err.(*exceptions.Exception); ok {
				c.JSON(ex.Status, ex)
				return
			}
			c.Error(err)

			bdrq := exceptions.BadRequest("Unknown error (500)", c)
			c.JSON(bdrq.Status, bdrq)
		}
	}
}
