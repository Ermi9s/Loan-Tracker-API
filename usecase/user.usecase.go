package usecase

import (
	"errors"

	"github.com/Ermi9s/Loan-Tracker-API/Loan-Tracker-API/domain"
)

type UserUsecase struct {
	UserRepository domain.User_Repository_interface
}

func NewUserUsecase(repository domain.User_Repository_interface) *UserUsecase {
	return &UserUsecase{UserRepository: repository}
}


func (usecase *UserUsecase) GetOneUser(id string) (domain.ResponseUser, error) {
	user, err := usecase.UserRepository.GetUserDocumentByID(id)
	if err != nil {
		return domain.ResponseUser{}, err
	}
	return domain.CreateResponseUser(user), nil
}

func (usecase *UserUsecase) GetUsers() ([]domain.ResponseUser, error) {
	users, err := usecase.UserRepository.GetUserDocuments()
	if err != nil {
		return []domain.ResponseUser{}, err
	}
	response_users := []domain.ResponseUser{}
	for _, user := range users {
		response_users = append(response_users, domain.CreateResponseUser(user))
	}
	return response_users, nil
}

func (usecase *UserUsecase) DeleteUser(id string) error {
	err := usecase.UserRepository.DeleteUserDocument(id)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *UserUsecase) FilterUser(filter map[string]string) ([]domain.ResponseUser, error) {
	users, err := usecase.UserRepository.FilterUserDocument(filter)
	if err != nil {
		return []domain.ResponseUser{}, err
	}
	response_users := []domain.ResponseUser{}
	for _, user := range users {
		response_users = append(response_users, domain.CreateResponseUser(user))
	}
	return response_users, nil
}

func (usecase *UserUsecase) UpdatePassword(id string, updated_user domain.UpdatePassword) (domain.ResponseUser, error) {
	if updated_user.Password != updated_user.Confirm {
		return domain.ResponseUser{}, errors.New("password and confirm password do not match")
	}

	user, err := usecase.UserRepository.UpdateUserPassword(id, updated_user.Password)
	if err != nil {
		return domain.ResponseUser{}, err
	}
	return domain.CreateResponseUser(user), nil
}
