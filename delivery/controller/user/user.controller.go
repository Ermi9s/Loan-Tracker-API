package controller

import (
	"github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	uc domain.User_Usecase_interface
}

func NewUserController(uc domain.User_Usecase_interface) *UserController {
	return &UserController{uc: uc}
}

func (controller *UserController)GetOneUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		iuser,_ := c.Get("user")
		luser := iuser.(domain.User)
		id := luser.ID.Hex()
		user, err := controller.uc.GetOneUser(id)
		if err != nil {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"user": user})
	}
}

func (controller *UserController)GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := controller.uc.GetUsers()
		if err != nil {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"users": users})
	}
}

func (controller *UserController)DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := controller.uc.DeleteUser(id)
		if err != nil {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"message": "user deleted"})
	}
}

func (controller *UserController)UpdatePassword() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var updated_user domain.UpdatePassword
		if err := c.ShouldBindJSON(&updated_user); err != nil {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}
		user, err := controller.uc.UpdatePassword(id, updated_user)
		if err != nil {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"user": user})
	}
}
