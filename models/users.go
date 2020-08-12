package models

import "time"

// UserModel
type UserModel struct {
	ID                       int64     `json:"id"`
	Name                     string    `json:"name"`
	Email                    string    `json:"email"`
	PasswordHash             string    `json:"-"`
	UserToken                string    `json:"user_token"`
	EmailConfirmationToken   string    `json:"email_confirmation_token"`
	EmailConfirmed           bool      `json:"email_confirmed"`
	ResetPasswordToken       string    `json:"reset_password_token"`
	ResetPasswordRequestedAt time.Time `json:"reset_password_requested_at"`
	CreatedAt                time.Time `json:"created_at"`
}

// GetEntityName()
func (UserModel) GetEntityName() string {
	return "users"
}

func (UserModel) TableName() string {
	return "users"
}

// UserLoginTokenModel
type UserLoginTokenModel struct {
	ID                      int64     `json:"id"`
	UserID                  int64     `json:"user_id"`
	UserLoginToken          string    `json:"user_login_token"`
	UserLoginTokenExpiresAt time.Time `json:"user_login_expires_at"`
}

func (UserLoginTokenModel) TableName() string {
	return "users_login_tokens"
}

// NewUser struct {
type NewUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoggedinResponse
type NewLoginResponse struct {
	Name                    string    `json:"name"`
	Email                   string    `json:"email"`
	CreatedAt               time.Time `json:"created_at"`
	UserLoginToken          string    `json:"user_login_token"`
	UserLoginTokenExpiresAt time.Time `json:"user_login_token_expires_at"`
}

// UserTokenResponse
type UserTokenResponse struct {
	UserToken string `json:"user_token"`
}

// ResetPasswordStart
type ResetPasswordStart struct {
	Email string `json:"email"`
}

// ResetPasswordConfirm
type ResetPasswordConfirm struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}
