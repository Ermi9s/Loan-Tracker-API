package repository

import (
	"context"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/database"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthRepository struct {
	Collection database.CollectionInterface
}


func NewAuthRepository(collection database.CollectionInterface) *AuthRepository {
	return &AuthRepository{Collection: collection}
}

func (authRepo *AuthRepository) RegisterUser(user domain.RegisterUser) (primitive.ObjectID , error) {
	useri := domain.User{
		UserName: user.UserName,
		Email: user.Email,
		Password: user.Password,
		Is_Admin: false,
	}
	inserted , err := authRepo.Collection.InsertOne(context.TODO() , useri)
	Id := inserted.InsertedID.(primitive.ObjectID)

	if err != nil {
		return primitive.NilObjectID , err
	}
	return Id , nil
}


func (authRepo *AuthRepository)InsertRefreshToken(id string  , token string) error {
	pid,err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	document := bson.D{{Key: "refresh_token" , Value: token}}

	filter := bson.D{{Key: "_id" , Value: pid}}
	update := bson.D{{Key: "$set" , Value: document}}

	updated_user := authRepo.Collection.FindOneAndUpdate(context.TODO() , filter , update)
	var new_user domain.User
	err  = updated_user.Decode(&new_user)
	if err != nil {
		return err
	}

	return nil
}


func (repository *AuthRepository)GetUserDocumentByEmail(email string) (domain.User , error) {
	filter := bson.D{{Key: "email" , Value: email}}
	result := repository.Collection.FindOne(context.TODO() , filter)

	var user domain.User
	err := result.Decode(&user)
	if err != nil {
		return domain.User{} , err
	}
	return user , nil
}
