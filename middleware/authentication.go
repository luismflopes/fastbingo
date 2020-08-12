package middleware

import (
	"fastbingo/models"
	"fastbingo/services"

	"github.com/gin-gonic/gin"
)

// AuthenticationMiddleware type description here
func AuthenticationMiddleware(us *services.UsersService) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-user-login-token")

		user := us.GetUserByLoginToken(token)
		if user == nil {
			panic(models.AppError{Code: models.Forbidden, Message: "User token is invalid or expired."})
		}

		c.Set("user", user)

		c.Next()
	}
}
