package systemlogger

import (
	"context"
	"time"

	"github.com/Loan-Tracker-API/Loan-Tracker-API/config"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/database"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
)

type Logger struct {
	logs_collection database.CollectionInterface
}


func (l *Logger) LogEvent(event string) error {
	connection  := config.ServerConnection{}
	connection.Connect_could()

	l.logs_collection = &database.MongoCollection{
		Collection: connection.Client.Database("LoanTracker").Collection("Logs"),
	}

	log := domain.Logs{
		Event: event,
		Time: time.Now().Weekday(),
	}

	_,err := l.logs_collection.InsertOne(context.TODO(), log)
	return err
}
