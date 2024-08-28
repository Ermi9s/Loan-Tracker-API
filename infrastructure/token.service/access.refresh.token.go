package tokenservice

import (
	"errors"
	"time"

	"github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
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


func (vt *TokenService_imp) GenrateToken(id string , expr int) (string, error) {
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

func (vt *TokenService_imp) GenrateRegistrationToken(user domain.RegisterUser) (string, error) {
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


func (vt *TokenService_imp) VerifyRegistrationToken(tokenStr string) (domain.RegisterUser , error) {
	token, err := jwt.ParseWithClaims(tokenStr, &domain.RegisterUserClaims{}, func(token *jwt.Token) (interface{}, error) {
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

