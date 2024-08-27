package tokenservice

import (
	"errors"
	"time"

	"github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/domain"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TokenService_imp struct {
	SecretKey []byte
}

func NewTokenService(Secret string) *TokenService_imp {
	return &TokenService_imp{
		SecretKey: []byte(Secret),
	}
}

func (t *TokenService_imp) GenerateAccessToken(id primitive.ObjectID) (string, error) {
	claims := domain.UserClaims{
        ID:      id,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
        },
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.SecretKey))
}

func (t *TokenService_imp) GenerateRefreshToken(id primitive.ObjectID) (string, error) {
	claims := domain.UserClaims{
        ID:      id,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour * 168).Unix(),
        },
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.SecretKey))
}

func (t *TokenService_imp) ValidateAccessToken(tokenStr string) (*domain.User, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid access token")
	}

	claims, ok := token.Claims.(*domain.UserClaims)
    if !ok {
        return nil, errors.New("invalid token claims")
    }

	return &domain.User{
        ID:      claims.ID,
    }, nil
}


func (t *TokenService_imp) ValidateRefreshToken(tokenStr string) (*domain.User, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(*domain.UserClaims)
    if !ok {
        return nil, errors.New("invalid token claims")
    }

	return &domain.User{
        ID:      claims.ID,
    }, nil
}

