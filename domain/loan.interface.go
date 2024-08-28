package domain

import "github.com/gin-gonic/gin"

type Loan_Repository_interface interface {
	GetAll(filter map[string]string) ([]Loan, error)
	GetLoanByID(id string) (Loan, error)
	CreateLoan(loan Loan) (Loan, error)
	DeleteLoan(id string) error
	UpdateLoan(id string, loan Loan) (Loan, error)
}

type Loan_Usecase_interface interface {
	GetAll(filter map[string]string) ([]Loan, error)
	GetLoanByID(id string) (Loan, error)
	CreateLoan(loan Loan, user ResponseUser) (Loan, error)
	DeleteLoan(id string) error
	UpdateLoan(id string, status string) (Loan, error)
}

type Loan_Controller_interface interface {
	GetAll() gin.HandlerFunc
	GetLoanByID() gin.HandlerFunc
	CreateLoan() gin.HandlerFunc
	DeleteLoan() gin.HandlerFunc
	UpdateLoan() gin.HandlerFunc
}

