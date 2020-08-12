package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"fastbingo/controllers"
	"fastbingo/middleware"
	"fastbingo/services"
	"fastbingo/storage"
)

func main() {
	db, err := gorm.Open("sqlite3", "store.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	storage.Migrate(db)

	ts := services.TranslationService{
		Storage: db,
	}

	// // instantiate controllers
	es := services.NewEmailService()

	us := services.UsersService{
		Storage:      db,
		EmailService: es,
	}

	uc := controllers.UsersController{
		UsersService: &us,
	}

	as := services.AuthService{
		UsersService: &us,
	}

	ac := controllers.AuthController{
		AuthService: &as,
	}

	wc := controllers.WelcomeController{}

	// define routes
	r := gin.New()
	r.Use(cors.Default())
	r.Use(gin.Logger())
	r.Use(middleware.LocaleMiddleware(&ts))
	r.Use(middleware.RecoveryMiddleware(middleware.RecoveryHandler))

	// Public routes
	r.GET("/", wc.WelcomeAction)
	r.GET("/api/v1", wc.WelcomeAction)

	r.POST("/api/v1/users", uc.CreateUserAction)
	r.GET("/api/v1/users/email/confirm/:ctoken", uc.ConfirmUserEmailAction)
	r.GET("/api/v1/users/email/resend/:utoken", uc.ResendUserConfirmationEmailAction)
	r.POST("/api/v1/users/password/reset", uc.ResetPasswordStartAction)
	r.PUT("/api/v1/users/password/reset", uc.ResetPasswordConfirmAction)

	r.POST("/api/v1/auth/login", ac.PasswordLoginAction)
	r.POST("/api/v1/auth/login/google", ac.GoogleLoginAction)
	r.POST("/api/v1/auth/login/facebook", ac.FacebookLoginAction)

	// User Routes
	loggedUser := r.Group("/api/v1")
	loggedUser.Use(middleware.AuthenticationMiddleware(&us))
	//loggedUser.GET("/users/me", uc.LoginUserAction)

	// Admin Routes

	// start the server
	r.Run("0.0.0.0:8081")
}
