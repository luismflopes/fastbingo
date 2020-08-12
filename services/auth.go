package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"fastbingo/models"

	"github.com/jinzhu/gorm"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// type ProductServiceContract interface {
// 	ListProduct() ([]*models.CmsPage, error)
// 	DetailProduct(slug string) (*models.CmsPage, error)
// 	CreateProduct(req *models.CmsPageCreateRequest)
// }

// AuthService
type AuthService struct {
	UsersService *UsersService
}

// PasswordLogin
func (me *AuthService) PasswordLogin(u *models.PasswordLoginRequest) *models.LoggedinResponse {

	var user models.UserModel
	err := me.UsersService.Storage.Where("email = ?", u.Email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			panic(models.AppError{Code: models.ErrorUserWrongCredentials})
		}
		panic(err)
	}

	if !user.EmailConfirmed {
		panic(models.AppError{Code: models.ErrorUserNotConfirmedEmail})
	}

	if !CheckPasswordHash(u.Password, user.PasswordHash) {
		panic(models.AppError{Code: models.ErrorUserWrongCredentials})
	}

	return me.UsersService.GenerateNewLoginToken(&user)
}

// GoogleLogin
func (me *AuthService) GoogleLogin(r *models.GoogleLoginRequest) *models.LoggedinResponse {

	conf := &oauth2.Config{
		ClientID:     "1083134170508-2l0smj8b6qn3bs3mh69c2cpo4d220jll",
		ClientSecret: "jYy0G_w0Qg3f8aJuepn328zj",
		RedirectURL:  "postmessage",
		Endpoint:     google.Endpoint,
	}

	// exchange auth_code to token including refresh_token
	// authCode is from frond-end side and only for one time use
	token, err := conf.Exchange(oauth2.NoContext, r.AuthCode)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
		return nil
	}

	u, err := fetchGoogleUserInfo(conf.Client(context.Background(), token))
	if err != nil {
		panic(models.AppError{Code: models.ErrorUnkown})
	}

	fmt.Println("Google parsed user:", u)

	return me.CheckUserAndGenerateNewLoginToken(u.Email, u.Name)
}

// FacebookLogin
func (me *AuthService) FacebookLogin(r *models.FacebookLoginRequest) *models.LoggedinResponse {
	u, err := fetchFacebookUserInfo(r.AccessToken)
	if err != nil {
		panic(models.AppError{Code: models.ErrorUnkown})
	}

	fmt.Println("Facebook parsed user:", u)

	return me.CheckUserAndGenerateNewLoginToken(u.Email, u.Name)
}

func (me *AuthService) CheckUserAndGenerateNewLoginToken(email, name string) *models.LoggedinResponse {
	user := me.UsersService.GetUserByEmail(email)
	if user == nil {
		user = &models.UserModel{
			Email:          email,
			Name:           name,
			PasswordHash:   HashPassword(GenerateRandomToken()),
			UserToken:      GenerateRandomToken(),
			EmailConfirmed: true,
		}

		err := me.UsersService.Storage.Create(&user).Error
		if err != nil {
			panic(err)
		}
	}

	return me.UsersService.GenerateNewLoginToken(user)
}

func fetchGoogleUserInfo(client *http.Client) (*models.GoogleUserInfo, error) {
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var gu models.GoogleUserInfo
	if err := json.Unmarshal(data, &gu); err != nil {
		return nil, err
	}
	return &gu, nil
}

func fetchFacebookUserInfo(accessToken string) (*models.FacebookUserInfo, error) {
	resp, err := http.Get("https://graph.facebook.com/me?fields=name,email&access_token=" +
		url.QueryEscape(accessToken))

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)
	}

	log.Printf("parseResponseBody: %s\n", string(data))

	var u models.FacebookUserInfo
	if err := json.Unmarshal(data, &u); err != nil {
		return nil, err
	}

	return &u, nil
}
