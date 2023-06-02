package service

import (
	"be17/main/feature/user"

	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userData user.UserDataInterface
	validate *validator.Validate
}

func New(repo user.UserDataInterface) user.UserServiceInterface {
	return &userService{
		userData: repo,
		validate: validator.New(),
	}
}

// Delete implements user.UserServiceInterface
func (service *userService) Delete(id int) error {
	err := service.userData.Delete(id)
	return err
}

// Update implements user.UserServiceInterface
func (service *userService) Update(id int, input user.Core) error {
	err := service.userData.Update(id, input)
	if err != nil {
		return fmt.Errorf("failed to update user with ID %d: %w", id, err)
	}
	return nil
}

// GetById implements user.UserServiceInterface
func (service *userService) GetById(id int) (user.Core, error) {
	data, err := service.userData.SelectById(id)
	return data, err
}

func (service *userService) Login(email string, password string) (user.Core, string, error) {
	if email == "" || password == "" {
		return user.Core{}, "", errors.New("error Validation: nama, email, password harus diisi ")
	}

	dataLogin, token, err := service.userData.Login(email, password)
	if err != nil {
		return user.Core{}, "", err
	}
	return dataLogin, token, nil
}

func (service *userService) Create(input user.Core) error {

	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	// Menyimpan data user ke database
	errInsert := service.userData.Insert(input)
	return errInsert
}
