package controller

import (
	"fmt"
	"github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/domain"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthUsecase domain.AuthUsecase_interface
}

func NewAuthController(usecase domain.AuthUsecase_interface) *AuthController {
	return &AuthController{AuthUsecase: usecase}
}

func (controller *AuthController) RegisterUser_Unverified() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user domain.RegisterUser
		err := c.BindJSON(&user)
		if err != nil {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}
		ruser , err := controller.AuthUsecase.RegisterUserU(user)
		if err != nil {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(
			200, 
			gin.H{"message":fmt.Sprintf("verification email sent to %s" , user.Email), 
			"user": ruser,
		})
}
} 

func (controller *AuthController) RegisterUser_verified() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Param("token")
		token , ruser , err := controller.AuthUsecase.RegisterUserV(token)
		if err != nil {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(
			200, 
			gin.H{"token":token, 
				"user": ruser,
				"message": "user registered successfully",
		})
}
} 

func (controller *AuthController) Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user domain.LogInUser
		err := c.BindJSON(&user)
		if err != nil {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}
		token , ruser , err := controller.AuthUsecase.LoginUser(user)
		if err != nil {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}

		c.IndentedJSON(
			200, 
			gin.H{"token":token, 
				"user": ruser,
				"message": "user logged in successfully",
		})
}
}