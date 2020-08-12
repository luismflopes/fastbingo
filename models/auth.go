package models

import "time"

// PasswordLoginRequest entity
type PasswordLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GoogleLoginRequest entity
type GoogleLoginRequest struct {
	AuthCode string `json:"auth_code"`
}

// FacebookLoginRequest entity
type FacebookLoginRequest struct {
	AccessToken string `json:"access_token"`
}

// LoggedinResponse
type LoggedinResponse struct {
	Name                    string    `json:"name"`
	Email                   string    `json:"email"`
	CreatedAt               time.Time `json:"created_at"`
	UserLoginToken          string    `json:"user_login_token"`
	UserLoginTokenExpiresAt time.Time `json:"user_login_token_expires_at"`
}

type GoogleUserInfo struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Profile       string `json:"profile"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Gender        string `json:"gender"`
}

type FacebookUserInfo struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
