package routes

import (
	"github.com/Loan-Tracker-API/Loan-Tracker-API/database"
	controller "github.com/Loan-Tracker-API/Loan-Tracker-API/delivery/controller/auth"

	"github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/email"
	enivronment "github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/env"
	passwordservice "github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/password"
	tokenservice "github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/token.service"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/repository"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/usecase"
	"github.com/gin-gonic/gin"
)

func NewAuthRoute(group *gin.RouterGroup, users database.CollectionInterface) {


	key := enivronment.OsGet("SECRETKEY")
	email_user := enivronment.OsGet("EMAIL_USER")
	email_password := enivronment.OsGet("EMAIL_PASSWORD")


	tkokenSev := tokenservice.NewTokenService(key)
	emailService := email.NewEmail(email_user, email_password)
	paswordServ := &passwordservice.PasswordS{}

	AuthRepository := repository.NewAuthRepository(users)
	Authusecase := usecase.NewAuthUsecase(AuthRepository , tkokenSev , paswordServ , emailService)
	AuthController := controller.NewAuthController(Authusecase)

	group.POST("/users/register", AuthController.RegisterUser_Unverified())
	group.GET("/users/verify-email/:token", AuthController.RegisterUser_verified())
	group.POST("/users/login", AuthController.Login())
	group.POST("/users/token/refresh" , AuthController.Refresh())
}