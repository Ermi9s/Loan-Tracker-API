package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type TokenSrvices interface {
	GenerateAccessToken(id primitive.ObjectID) (string, error)
	GenerateRefreshToken(id primitive.ObjectID) (string, error)
	ValidateAccessToken(tokenStr string) (*User, error)
	ValidateRefreshToken(tokenStr string) (*User, error)
	GenrateToken(id string, expr int)
	VerifyRegistrationToken(tokenStr string) (RegisterUser , error)
	GenrateRegistrationToken(user RegisterUser) (string, error)
}
