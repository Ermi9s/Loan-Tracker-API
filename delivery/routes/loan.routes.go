package routes

import (
	"github.com/Loan-Tracker-API/Loan-Tracker-API/database"
	controller "github.com/Loan-Tracker-API/Loan-Tracker-API/delivery/controller/loan"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/repository"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/usecase"
	"github.com/gin-gonic/gin"
)

func NewLoanRoutes(group gin.RouterGroup , loan_collection database.CollectionInterface) {
	loan_repository := repository.NewLoanRepository(loan_collection)
	loan_usecase := usecase.NewLoanUsecase(loan_repository)
	loan_controller := controller.NewLoanController(loan_usecase)

	group.GET("/admin/loans", loan_controller.GetAll())
	group.GET("/loans/:id", loan_controller.GetLoanByID())
	group.POST("/loans", loan_controller.CreateLoan())
	group.DELETE("/loans/:id", loan_controller.DeleteLoan())
	group.PATCH("/loans/:id/:status", loan_controller.UpdateLoan())
}