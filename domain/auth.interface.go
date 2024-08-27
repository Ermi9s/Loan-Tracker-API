package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type AuthRepository_interface interface {
	RegisterUser(user RegisterUser) (primitive.ObjectID, error)
	InsertRefreshToken(id string  , token string) error
	GetUserDocumentByEmail(email string) (User , error)
}

type PasswordServices interface {
	HashPassword(password string) (string, error)
	ComparePassword(DBpassword string, InputPassword string) (bool, error)
}
type AuthUsecase_interface interface {
	RegisterUserV(token string) (string , ResponseUser , error)
	RegisterUserU(user RegisterUser) (ResponseUser , error)
	LoginUser(user LogInUser) (string , ResponseUser , error)
	RefreshTokens(refreshToken string) (string, string, error)
}

type EmailServices interface {
	SendVerificationEmail(to, subject, body string) error
}


type EmailVUsecase interface {
	SendForgretPasswordEmail(id string, vuser VerifyUser) error
	ValidateForgetPassword(id string, token string) error
	UpdatePassword(id string, update_password UpdatePassword) (User, error)
}