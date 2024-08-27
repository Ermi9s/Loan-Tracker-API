package passwordservice

import (
	"errors"	
	domain "github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
	enivronment "github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/env"
	"github.com/dgrijalva/jwt-go"
)

func IsValidForgetToken(strtoken string, id string) error {

	var SecretKey = []byte(enivronment.OsGet("SECRETKEY"))

	token, err := jwt.ParseWithClaims(strtoken, &domain.UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {

		return errors.New("token not valid")
	}

	payload, ok := token.Claims.(*domain.UserClaims)
	if !ok {
		return errors.New("token payload not valid")
	}

	if id != payload.ID.Hex() {
		return errors.New("user-Id not matching the token-id")
	}

	return nil
}