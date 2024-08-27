package repository

import (
	"context"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/database"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanRepository struct {
	loan_collection database.CollectionInterface
}

func NewLoanRepository(loan_collection database.CollectionInterface) *LoanRepository {
	return &LoanRepository{
		loan_collection: loan_collection,
	}
}

func (lr *LoanRepository) GetAll(filter map[string]string) ([]domain.Loan, error) {
	bfilter := bson.D{}

	for key, value := range filter {
		bfilter = append(bfilter, bson.E{Key: key, Value: value})
	}
	loans, err := lr.loan_collection.Find(context.TODO() ,bfilter)
	if err != nil {
		return nil, err
	}

	rloans := []domain.Loan{}
	for loans.Next(context.TODO()){
		var loan domain.Loan
		err := loans.Decode(&loan)
		if err != nil {
			return nil, err
		}
		rloans = append(rloans, loan)
	}

	return rloans, nil
}

func (lr *LoanRepository) GetLoanByID(id string) (domain.Loan, error) {
	objId,_ := primitive.ObjectIDFromHex(id)
	bfilter := bson.D{{Key: "id", Value: objId}}
	loan := domain.Loan{}
	err := lr.loan_collection.FindOne(context.TODO(), bfilter).Decode(&loan)
	if err != nil {
		return domain.Loan{}, err
	}

	return loan, nil
}

func (lr *LoanRepository) CreateLoan(loan domain.Loan) (domain.Loan, error) {
	_, err := lr.loan_collection.InsertOne(context.TODO(), loan)
	if err != nil {
		return domain.Loan{}, err
	}
	return loan, nil
}

func (lr *LoanRepository) DeleteLoan(id string) error {
	objId,_ := primitive.ObjectIDFromHex(id)
	bfilter := bson.D{{Key: "id", Value: objId}}
	_, err := lr.loan_collection.DeleteOne(context.TODO(), bfilter)
	if err != nil {
		return err
	}
	return nil
}

func (lr *LoanRepository) UpdateLoan(id string, loan domain.Loan) (domain.Loan, error) {
	objId,_ := primitive.ObjectIDFromHex(id)
	bfilter := bson.D{{Key: "id", Value: objId}}
	_,err := lr.loan_collection.UpdateOne(context.TODO(), bfilter, loan)
	if err != nil {
		return domain.Loan{}, err
	}
	return loan, nil
}