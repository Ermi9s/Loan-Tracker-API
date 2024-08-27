package domain

import (
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	Is_Admin     bool               `json:"is_admin" bson:"is_admin"`
	Password     string             `json:"password,omitempty" bson:"password,omitempty"`
	RefreshToken string             `json:"refresh_token,omitempty" bson:"refresh_token,omitempty"`
}

type ResponseUser struct {
	ID       string `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}

type LogInUser struct {
	UserName string `json:"username,omitempty" bson:"username,omitempty"`
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password" bson:"password"`
}

type RegisterUser struct {
	UserName string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type UserClaims struct {
	jwt.StandardClaims
	ID primitive.ObjectID `json:"_id,"`
}

type UpdatePassword struct {
	Password string `json:"password" bson:"password"`
	Confirm  string `json:"confirm" bson:"confirm"`
}


type RegisterUserClaims struct {
	jwt.StandardClaims
	UserName     string             `json:"username" bson:"username"`
	Email        string             `json:"email" bson:"email"`
	Password     string             `json:"password"`
}

// from actual user model to response model to be done in usecase
func CreateResponseUser(user User) ResponseUser {
	return ResponseUser{
		ID:       user.ID.Hex(),
		UserName: user.UserName,
		Email:    user.Email,
	}
}


