package middleware

import (
	"fastbingo/services"

	"github.com/gin-gonic/gin"
)

// LocaleMiddleware type description here
func LocaleMiddleware(ts *services.TranslationService) gin.HandlerFunc {
	return func(c *gin.Context) {
		l, err := ts.CheckLocale(c.Request.Header.Get("accept-language"))
		if err != nil {
			l = "en"
		}
		c.Request.Header.Set("accept-language", l)

		c.Next()
	}
}
