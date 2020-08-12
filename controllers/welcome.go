package controllers

import (
	"github.com/gin-gonic/gin"
)

type WelcomeController struct {
}

func (me *WelcomeController) WelcomeAction(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 0,
		"message": "FastBingo Api ready!",
	})
}
