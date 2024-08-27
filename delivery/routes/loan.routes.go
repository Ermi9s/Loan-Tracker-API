package routes

import (
	"github.com/Loan-Tracker-API/Loan-Tracker-API/database"
	controller "github.com/Loan-Tracker-API/Loan-Tracker-API/delivery/controller/loan"
	enivronment "github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/env"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/middleware"
	tokenservice "github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/token.service"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/repository"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/usecase"
	"github.com/gin-gonic/gin"
)

func NewLoanRoutes(group *gin.RouterGroup , loan_collection database.CollectionInterface ,  user_collection database.CollectionInterface) {
	loan_repository := repository.NewLoanRepository(loan_collection)
	loan_usecase := usecase.NewLoanUsecase(loan_repository)
	loan_controller := controller.NewLoanController(loan_usecase)
	
	key := enivronment.OsGet("SECRETKEY")
	tkokenSev := tokenservice.NewTokenService(key)

	loggedInMiddleware := middleware.LoggedIn(*tkokenSev)
	mustOwnMiddleware := middleware.RoleBasedAuth(false ,user_collection )
	mustBeAdminMiddleware := middleware.RoleBasedAuth(true , user_collection)

	group.GET("/admin/loans", loggedInMiddleware , mustBeAdminMiddleware , loan_controller.GetAll())
	group.PATCH("/loans/:id/:status",loggedInMiddleware,mustBeAdminMiddleware,loan_controller.UpdateLoan())

	group.GET("/loans/:id",loggedInMiddleware , mustOwnMiddleware,  loan_controller.GetLoanByID())

	group.POST("/loans", loggedInMiddleware , loan_controller.CreateLoan())
	group.DELETE("/loans/:id", loggedInMiddleware , loan_controller.DeleteLoan())
}