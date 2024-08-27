package tokenservice

import (
	"errors"
	"time"

	"github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/domain"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type VerifyToken struct{
	SecretKey []byte
}

func NewVerifyToken(Secret string) *VerifyToken {
	return &VerifyToken{
		SecretKey: []byte(Secret),
	}
}

func (vt *VerifyToken) GenrateToken(id string , expr int) (string, error) {
	obJID,_ := primitive.ObjectIDFromHex(id)
	itoken := jwt.NewWithClaims(jwt.SigningMethodHS256, domain.UserClaims{
		ID:    obJID,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour *time.Duration(expr)).Unix()},
	})
	token, err := itoken.SignedString(vt.SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (vt *VerifyToken) GenrateRegistrationToken(user domain.RegisterUser) (string, error) {
	claims := domain.RegisterUserClaims{
		UserName:     user.UserName,
		Email:        user.Email,
		Password:     user.Password,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Hour *1).Unix()},
	}

	itoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := itoken.SignedString(vt.SecretKey)
	if err != nil {
		return "", err
	}
	return token, nil
}


func (vt *VerifyToken) VerifyRegistrationToken(tokenStr string) (domain.RegisterUser , error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(vt.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return domain.RegisterUser{}, errors.New("invalid access token")
	}

	claims, ok := token.Claims.(*domain.RegisterUserClaims)
    if !ok {
        return domain.RegisterUser{}, errors.New("invalid token claims")
    }

	return domain.RegisterUser{
        UserName: claims.UserName,
		Email:   claims.Email,
		Password: claims.Password,
    }, nil
}
