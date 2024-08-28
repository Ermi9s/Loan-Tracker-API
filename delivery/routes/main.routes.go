package routes

import (
	"github.com/Loan-Tracker-API/Loan-Tracker-API/config"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/database"
	"github.com/gin-gonic/gin"
)

func NewMainRoute() {
	router := gin.Default()

	mogodb := config.ServerConnection{}
	mogodb.Connect_could()
	mongo_database := mogodb.Client.Database("LoanTracker")

	users := &database.MongoCollection{
		Collection: mongo_database.Collection("Users"),
	}

	loan := &database.MongoCollection{
		Collection: mongo_database.Collection("Loans"),
	}
	
	authRoute := router.Group("")
	loanRoute := router.Group("")
	resetPasswordRoute := router.Group("")

	NewAuthRoute(authRoute , users)
	NewLoanRoutes(loanRoute , loan , users)
	NewPasswordResetRoute(resetPasswordRoute , users)

	router.Run(":8080")
}
