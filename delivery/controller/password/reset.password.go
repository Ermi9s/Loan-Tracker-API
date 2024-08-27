package controller

import (
	"fmt"
	"net/http"
	domain "github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/domain"
	"github.com/gin-gonic/gin"
)

type EmailVControler struct {
	email_uc domain.EmailVUsecase
}

func NewEmailVControler(email_uc domain.EmailVUsecase) *EmailVControler {
	return &EmailVControler{email_uc: email_uc}
}

func (ctrl *EmailVControler) ForgetPasswordValidate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		token := ctx.Query("token")

		err := ctrl.email_uc.ValidateForgetPassword(id, token)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		update_password := domain.UpdatePassword{
			Password:  "12345678",
			Confirm: "12345678",
		}
		user, err := ctrl.email_uc.UpdatePassword(id, update_password)
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.IndentedJSON(http.StatusAccepted, gin.H{"user": user, "message": "your password is reset to 12345678, you can change it anytime you want"})
	}
}

func (ctrl *EmailVControler) SendForgetPasswordEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var model domain.VerifyUser
		if err := ctx.BindJSON(&model); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := ctx.Param("id")
		if err := ctrl.email_uc.SendForgretPasswordEmail(id, model); err != nil {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errorss": err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusAccepted, gin.H{"message": fmt.Sprintf("email sent to: %s", model.Email)})
	}
}
