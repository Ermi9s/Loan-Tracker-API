package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Loan struct {
	ID     primitive.ObjectID `json:"_id" bson:"_id"`
	Amount float64            `json:"amount" bson:"amount"`
	Status string             `json:"status" bson:"status"`
	OwedBy string             `json:"owedBy,omitempty" bson:"owedBy,omitempty"`
}
