package controllers

import (
	"encoding/json"
	"fmt"
	"fastbingo/models"
	"fastbingo/services"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// AuthController
type AuthController struct {
	AuthService *services.AuthService
}

// LoginUserAction
func (me *AuthController) PasswordLoginAction(c *gin.Context) {

	fmt.Println(c.Get("user"))

	var u *models.PasswordLoginRequest
	json.NewDecoder(c.Request.Body).Decode(&u)
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	l := me.AuthService.PasswordLogin(u)
	c.JSON(200, l)
	return
}

// GoogleLoginAction
func (me *AuthController) GoogleLoginAction(c *gin.Context) {
	var r *models.GoogleLoginRequest
	json.NewDecoder(c.Request.Body).Decode(&r)
	_, err := govalidator.ValidateStruct(r)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	l := me.AuthService.GoogleLogin(r)
	c.JSON(200, l)
	return
}

// FacebookLoginAction
func (me *AuthController) FacebookLoginAction(c *gin.Context) {
	var r *models.FacebookLoginRequest
	json.NewDecoder(c.Request.Body).Decode(&r)
	_, err := govalidator.ValidateStruct(r)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	l := me.AuthService.FacebookLogin(r)
	c.JSON(200, l)
	return
}
