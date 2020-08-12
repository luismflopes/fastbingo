package middleware

import (
	"fmt"
	"runtime/debug"
	"fastbingo/models"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// RecoveryMiddleware Last level to handle a panic
func RecoveryMiddleware(f func(c *gin.Context, err *models.AppError)) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("RecoveryMiddleware: ", err)
				debug.PrintStack()
				e := models.AppError{Code: models.ErrorUnkown, Message: models.ErrorUnkownMessage}

				switch v := err.(type) {
				case models.AppError:
					e = v
				case string:
					e.Message = v
				}

				if err == gorm.ErrRecordNotFound {
					e = models.AppError{Code: models.ErrorNotFound, Message: models.ErrorNotFound}
				}

				f(c, &e)
			}
		}()
		c.Next()
	}
}

func RecoveryHandler(c *gin.Context, err *models.AppError) {
	c.JSON(err.GetErrorStatusCode(), err)
}
