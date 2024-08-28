package controller

import (
	"github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
	"github.com/gin-gonic/gin"
)

type LoanController struct {
	loan_usecase domain.Loan_Usecase_interface
}

func NewLoanController(loan_usecase domain.Loan_Usecase_interface) *LoanController {
	return &LoanController{
		loan_usecase: loan_usecase,
	}
}


func (lc *LoanController) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		filter := make(map[string]string)
		for key, value := range c.Request.URL.Query() {
			filter[key] = value[0]
		}
		loans, err := lc.loan_usecase.GetAll(filter)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, loans)
	}
}

func (lc *LoanController) GetLoanByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id  := c.Param("id")
	
		loan, err := lc.loan_usecase.GetLoanByID(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, loan)
	}
}

func (lc *LoanController) CreateLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loan domain.Loan
		err := c.BindJSON(&loan)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		iuser ,_ := c.Get("user")
		user := iuser.(domain.ResponseUser)

		loan, err = lc.loan_usecase.CreateLoan(loan, user)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, loan)
	}
}

func (lc *LoanController) DeleteLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		id  := c.Param("id")
	
		err := lc.loan_usecase.DeleteLoan(id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "Loan Deleted"})
	}
}

func (lc *LoanController) UpdateLoan() gin.HandlerFunc {
	return func(c *gin.Context) {
		id  := c.Param("id")
		status  := c.Param("status")
	
		loan, err := lc.loan_usecase.UpdateLoan(id, status)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, loan)
	}
}

