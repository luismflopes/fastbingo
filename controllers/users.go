package controllers

import (
	"encoding/json"
	"fastbingo/models"
	"fastbingo/services"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

// UsersController
type UsersController struct {
	UsersService *services.UsersService
}

// CreateUserAction
func (me *UsersController) CreateUserAction(c *gin.Context) {
	var u *models.NewUser
	json.NewDecoder(c.Request.Body).Decode(&u)
	_, err := govalidator.ValidateStruct(u)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	user := me.UsersService.CreateUser(u)

	c.JSON(200, user)
	return
}

// ConfirmUserEmailAction
func (me *UsersController) ConfirmUserEmailAction(c *gin.Context) {

	token, _ := c.Params.Get("ctoken")

	l := me.UsersService.ConfirmUserEmail(token)
	c.JSON(200, l)
	return
}

// ResendUserConfirmationEmailAction
func (me *UsersController) ResendUserConfirmationEmailAction(c *gin.Context) {
	token, _ := c.Params.Get("utoken")

	user := me.UsersService.ResendUserConfirmationEmail(token)
	c.JSON(200, user)
	return
}

func (me *UsersController) ResetPasswordStartAction(c *gin.Context) {
	var r *models.ResetPasswordStart
	json.NewDecoder(c.Request.Body).Decode(&r)
	_, err := govalidator.ValidateStruct(r)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	user := me.UsersService.ResetPasswordStart(r)
	c.JSON(200, user)
	return
}

func (me *UsersController) ResetPasswordConfirmAction(c *gin.Context) {
	var r *models.ResetPasswordConfirm
	json.NewDecoder(c.Request.Body).Decode(&r)
	_, err := govalidator.ValidateStruct(r)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	user := me.UsersService.ResetPasswordConfirm(r)
	c.JSON(200, user)
	return
}
