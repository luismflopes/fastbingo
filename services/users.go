package services

import (
	"fastbingo/models"
	"time"

	"github.com/jinzhu/gorm"
)

// type ProductServiceContract interface {
// 	ListProduct() ([]*models.CmsPage, error)
// 	DetailProduct(slug string) (*models.CmsPage, error)
// 	CreateProduct(req *models.CmsPageCreateRequest)
// }

// UsersService
type UsersService struct {
	Storage      *gorm.DB
	EmailService *EmailService
}

// GetUserByEmail
func (me *UsersService) GetUserByEmail(email string) *models.UserModel {

	var user models.UserModel
	err := me.Storage.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		panic(err)
	}

	return &user
}

// GetUserByLoginToken
func (me *UsersService) GetUserByLoginToken(token string) *models.UserLoginTokenModel {
	u := models.UserLoginTokenModel{}

	query := `select u.*
	from users_login_tokens ut
	inner join users u ON u.id=ut.user_id
	where ut.token=? AND ut.token_expires_at > date('now')`

	err := me.Storage.Raw(query, token).Scan(&u).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil
		}

		panic(err.Error())
	}

	return &u
}

// CreateUser
func (me *UsersService) CreateUser(u *models.NewUser) *models.UserTokenResponse {

	user := models.UserModel{
		Name:                   u.Name,
		Email:                  u.Email,
		PasswordHash:           HashPassword(u.Password),
		UserToken:              GenerateRandomToken(),
		EmailConfirmationToken: GenerateRandomToken(),
	}

	err := me.Storage.Create(&user).Error
	if err != nil {
		panic(err)
	}

	me.EmailService.SendEmailNewAccount(user.Email, user.Name, user.EmailConfirmationToken)

	return &models.UserTokenResponse{
		UserToken: user.UserToken,
	}
}

// ConfirmUserEmail
func (me *UsersService) ConfirmUserEmail(confirmToken string) *models.LoggedinResponse {

	var user models.UserModel
	err := me.Storage.Where("email_confirmation_token = ?", confirmToken).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(models.AppError{Code: models.ErrorNotFound})
		}
		panic(err)
	}

	err = me.Storage.Model(&user).Updates(map[string]interface{}{
		"email_confirmation_token": "",
		"email_confirmed":          true,
	}).Error
	if err != nil {
		panic(err)
	}

	return me.GenerateNewLoginToken(&user)
}

// ResendUserConfirmationEmail
func (me *UsersService) ResendUserConfirmationEmail(userToken string) *models.UserTokenResponse {

	var user *models.UserModel
	err := me.Storage.Where("user_token = ?", userToken).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(models.AppError{Code: models.ErrorNotFound})
		}
		panic(err)
	}

	err = me.Storage.Model(&user).Update("email_confirmation_token", GenerateRandomToken()).Error
	if err != nil {
		panic(err)
	}

	go me.EmailService.SendEmailNewAccount(user.Email, user.Name, user.EmailConfirmationToken)

	return &models.UserTokenResponse{
		UserToken: user.UserToken,
	}
}

// ResetPasswordStart
func (me *UsersService) ResetPasswordStart(rp *models.ResetPasswordStart) *models.UserTokenResponse {
	var user models.UserModel
	err := me.Storage.Where("email = ?", rp.Email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(models.AppError{Code: models.ErrorNotFound})
		}
		panic(err)
	}

	if time.Now().Unix() < user.ResetPasswordRequestedAt.Add(time.Minute*1).Unix() {
		panic(models.AppError{Code: models.ErrorBadRequest, Message: "A new request within one minute"})
	}

	err = me.Storage.Model(&user).Update(map[string]interface{}{
		"reset_password_token":        GenerateRandomToken(),
		"reset_password_requested_at": time.Now(),
	}).Error
	if err != nil {
		panic(err)
	}

	go me.EmailService.SendEmailResetPassword(
		user.Email,
		user.Name,
		user.ResetPasswordToken)

	return &models.UserTokenResponse{
		UserToken: user.UserToken,
	}
}

// ResetPasswordConfirm
func (me *UsersService) ResetPasswordConfirm(r *models.ResetPasswordConfirm) *models.LoggedinResponse {
	var user models.UserModel
	err := me.Storage.Where("email = ?", "luismiguelflo@gmail.com").First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(models.AppError{Code: models.ErrorNotFound})
		}
		panic(err)
	}

	if false {
		if time.Now().Unix() > user.ResetPasswordRequestedAt.Add(time.Hour*1).Unix() {
			panic(models.AppError{Code: models.ErrorBadRequest, Message: "Reset Password Token expired"})
		}

		// clear reset token and set new password
		err = me.Storage.Model(&user).Update(map[string]interface{}{
			"reset_password_token":        gorm.Expr("NULL"),
			"reset_password_requested_at": gorm.Expr("NULL"),
			"password_hash":               HashPassword(r.Password),
		}).Error
		if err != nil {
			panic(err)
		}

		err = me.Storage.Delete(models.UserLoginTokenModel{}, "user_id=?", user.ID).Error
		if err != nil {
			panic(err)
		}
	}

	return me.GenerateNewLoginToken(&user)
}

// GenerateNewLoginToken
func (me *UsersService) GenerateNewLoginToken(user *models.UserModel) *models.LoggedinResponse {
	token := models.UserLoginTokenModel{
		UserID:                  user.ID,
		UserLoginToken:          GenerateRandomToken(),
		UserLoginTokenExpiresAt: time.Now().AddDate(0, 1, 1).UTC(),
	}

	err := me.Storage.Create(&token).Error
	if err != nil {
		panic(err)
	}

	return &models.LoggedinResponse{
		Name:                    user.Name,
		Email:                   user.Email,
		CreatedAt:               user.CreatedAt,
		UserLoginToken:          token.UserLoginToken,
		UserLoginTokenExpiresAt: token.UserLoginTokenExpiresAt,
	}
}
