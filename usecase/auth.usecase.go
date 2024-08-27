package usecase

import (
	"github.com/Loan-Tracker-API/Loan-Tracker-API/config"
	domain "github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
)

type AuthUsecase struct {
	authRepo     domain.AuthRepository_interface
	tokenServ    domain.TokenSrvices
	passwordServ domain.PasswordServices
	emailServ    domain.EmailServices
}

func NewAuthUsecase(
	repository domain.AuthRepository_interface,
	tokenServ domain.TokenSrvices,
	passwordServ domain.PasswordServices,
	emailServ domain.EmailServices,
) *AuthUsecase {
	return &AuthUsecase{
		authRepo: repository}
}

func (usecase *AuthUsecase) RegisterUserV(token string) (string, domain.ResponseUser, error) {
	user, err := usecase.tokenServ.VerifyRegistrationToken(token)
	if err != nil {
		return "", domain.ResponseUser{}, err
	}

	hashed, err := usecase.passwordServ.HashPassword(user.Password)
	if err != nil {
		return "", domain.ResponseUser{}, err
	}
	user.Password = hashed
	id, err := usecase.authRepo.RegisterUser(user)
	if err != nil {
		return "", domain.ResponseUser{}, err
	}
	Atoken, err := usecase.tokenServ.GenerateAccessToken(id)
	if err != nil {
		return "", domain.ResponseUser{}, err
	}

	Rtoken, err := usecase.tokenServ.GenerateRefreshToken(id)
	if err != nil {
		return "", domain.ResponseUser{}, err
	}

	err = usecase.authRepo.InsertRefreshToken(id.Hex(), Rtoken)
	if err != nil {
		return "", domain.ResponseUser{}, err
	}

	return Atoken, domain.ResponseUser{ID: id.Hex(), UserName: user.UserName, Email: user.Email}, nil
}

func (usecase *AuthUsecase) RegisterUserU(user domain.RegisterUser) (domain.ResponseUser, error) {
	token, err := usecase.tokenServ.GenrateRegistrationToken(user)
	if err != nil {
		return domain.ResponseUser{}, err
	}

	body, sub := config.ConfigBody(token)
	err = usecase.emailServ.SendVerificationEmail(user.Email, sub, body)
	if err != nil {
		return domain.ResponseUser{}, err
	}

	return domain.ResponseUser{UserName: user.UserName, Email: user.Email}, nil
}

func (usecase *AuthUsecase) LoginUser(user domain.LogInUser) (string, domain.ResponseUser, error) {
	userDoc, err := usecase.authRepo.GetUserDocumentByEmail(user.Email)
	if err != nil {
		return "", domain.ResponseUser{}, err
	}

	match, err := usecase.passwordServ.ComparePassword(userDoc.Password, user.Password)
	if err != nil {
		return "", domain.ResponseUser{}, err
	}
	if !match {
		return "", domain.ResponseUser{}, err
	}

	Atoken, err := usecase.tokenServ.GenerateAccessToken(userDoc.ID)
	if err != nil {
		return "", domain.ResponseUser{}, err
	}

	Rtoken, err := usecase.tokenServ.GenerateRefreshToken(userDoc.ID)
	if err != nil {
		return "", domain.ResponseUser{}, err
	}

	err = usecase.authRepo.InsertRefreshToken(userDoc.ID.Hex(), Rtoken)
	if err != nil {
		return "", domain.ResponseUser{}, err
	}

	return Atoken, domain.ResponseUser{ID: userDoc.ID.Hex(), UserName: userDoc.UserName, Email: userDoc.Email}, nil
}

func (u *AuthUsecase) RefreshTokens(refreshToken string) (string, string, error) {
    user, err := u.tokenServ.ValidateRefreshToken(refreshToken)
    if err != nil {
        return "", "", err
    }

    newAccessToken, err := u.tokenServ.GenerateAccessToken(user.ID)
    if err != nil {
        return "", "", err
    }

    newRefreshToken, err := u.tokenServ.GenerateRefreshToken(user.ID)
    if err != nil {
        return "", "", err
    }

    return newAccessToken, newRefreshToken, nil
}
