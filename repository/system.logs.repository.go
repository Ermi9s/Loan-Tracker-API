package repository

import (
	"context"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/database"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
	"go.mongodb.org/mongo-driver/bson"
)

type SystemLogsRepository struct {
	user_collection database.CollectionInterface
	loan_collection database.CollectionInterface
	log_collection database.CollectionInterface
}


func NewSystemLogsRepository(user_collection database.CollectionInterface, loan_collection database.CollectionInterface) *SystemLogsRepository {
	return &SystemLogsRepository{
		user_collection: user_collection,
		loan_collection: loan_collection,
	}
}


func (slr *SystemLogsRepository) GetAllEvents() ([]domain.Logs, error) {
	var logs []domain.Logs
	cursor, err := slr.log_collection.Find(context.TODO() , bson.D{})
	if err != nil {
		return logs, err
	}
	for cursor.Next(context.TODO()) {
		var log domain.Logs
		err := cursor.Decode(&log)
		if err != nil {
			return logs, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}