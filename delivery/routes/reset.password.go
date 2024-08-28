package routes
import (
	"github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/database"
	controller "github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/delivery/controller/password"
	"github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/infrastructure/email"
	enivronment "github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/infrastructure/env"
	passwordservice "github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/infrastructure/password"
	tokenservice "github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/infrastructure/token.service"
	"github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/repository"
	"github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/usecase"
	"github.com/gin-gonic/gin"
)


func NewVerifyEmialRoute(group *gin.RouterGroup, user_collection database.CollectionInterface) {
	key := enivronment.OsGet("SECRETKEY")
	email_user := enivronment.OsGet("EMAIL_USER")
	email_password := enivronment.OsGet("EMAIL_PASSWORD")
	
	repo := repository.NewUserRepository(user_collection)
	tkokenSev := tokenservice.NewTokenService(key)
	emailService := email.NewEmail(email_user, email_password)
	paswordServ := &passwordservice.PasswordS{}

	uc := usecase.NewEmailVUsecase(emailService, tkokenSev, repo, paswordServ)
	ctrl := controller.NewEmailVControler(uc)

	group.POST("/users/password-reset", ctrl.SendForgetPasswordEmail())
	group.GET("/users/password-update", ctrl.ForgetPasswordValidate())
}