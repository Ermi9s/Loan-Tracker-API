package usecase

import (
	"github.com/Loan-Tracker-API/Loan-Tracker-API/config"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
	passwordservice "github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/password"
)

type EmailVUsecase struct {
	repo        domain.User_Repository_interface
	emailSrv    domain.EmailServices
	tokenSrv    domain.TokenSrvices
	passwordSrv domain.PasswordServices
}

func NewEmailVUsecase(emailSrv domain.EmailServices,
	tokenSrv domain.TokenSrvices,
	repo domain.User_Repository_interface,
	passwordSrv domain.PasswordServices) *EmailVUsecase {
	return &EmailVUsecase{
		emailSrv: emailSrv,
		tokenSrv: tokenSrv,
		repo:    repo,
		passwordSrv: passwordSrv,
	}
}

func (uc *EmailVUsecase) SendForgretPasswordEmail(id string, vuser domain.VerifyUser) error {
	token, err := uc.tokenSrv.GenrateToken(id, 1)
	if err != nil {
		return err
	}
	subject, body := config.ConfigFogetBody(token, id)

	err = uc.emailSrv.SendVerificationEmail(vuser.Email, subject, body)
	if err != nil {
		return err
	}

	return nil
}

func (uc *EmailVUsecase) ValidateForgetPassword(id string, token string) error {
	return passwordservice.IsValidForgetToken(token, id)
}

func (uc *EmailVUsecase) UpdatePassword(id string, update_password domain.UpdatePassword) (domain.User, error) {
	hashed_password, err := uc.passwordSrv.HashPassword(update_password.Password)
	if err != nil {
		return domain.User{}, err
	}
	user, err := uc.repo.UpdateUserPassword(id, hashed_password)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
