package usecase

import (
	domain "github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
)

type LoanUsecase struct {
	loan_repository domain.Loan_Repository_interface
}

func NewLoanUsecase(loan_repository domain.Loan_Repository_interface) *LoanUsecase {
	return &LoanUsecase{
		loan_repository: loan_repository,
	}
}

func (lu *LoanUsecase) GetAll(filter map[string]string) ([]domain.Loan, error) {
	return lu.loan_repository.GetAll(filter)
}

func (lu *LoanUsecase) GetLoanByID(id string) (domain.Loan, error) {
	return lu.loan_repository.GetLoanByID(id)
}

func (lu *LoanUsecase) CreateLoan(loan domain.Loan , user domain.ResponseUser) (domain.Loan, error) {
	loan.OwedBy = user
	return lu.loan_repository.CreateLoan(loan)
}

func (lu *LoanUsecase) DeleteLoan(id string) error {
	return lu.loan_repository.DeleteLoan(id)
}

func (lu *LoanUsecase) UpdateLoan(id string, status string) (domain.Loan, error) {
	loan, err := lu.loan_repository.GetLoanByID(id)
	if err != nil {
		return domain.Loan{}, err
	}
	loan.Status = status
	return lu.loan_repository.UpdateLoan(id, loan)
}