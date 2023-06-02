package data

import (
	"be17/main/app/middlewares"
	"be17/main/feature/user"
	"be17/main/helper"
	"errors"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserDataInterface {
	return &userQuery{
		db: db,
	}
}


// Login implements user.UserDataInterface
func (repo *userQuery) Login(email string, password string) (user.Core, string, error) {
	var userGorm User

	tx := repo.db.Where("email=?", email).First(&userGorm)
	if tx.Error != nil {
		return user.Core{}, "", tx.Error
	}
	if tx.RowsAffected == 0 {
		return user.Core{}, "", errors.New("login failed, email dan password")
	}
	checkPassword := helper.CheckPasswordHash(password, userGorm.Password)
	if !checkPassword {
		return user.Core{}, "", errors.New("login failed, password salah")
	}

	token, errToken := middlewares.CreateToken(int(userGorm.ID))
	if errToken != nil {
		return user.Core{}, "", errToken
	}
	dataCore := ModelToCore(userGorm)

	return dataCore, token, nil
}

// insert implements user.UserDataInterface

func (repo *userQuery) Insert(input user.Core) error {
	hashedPassword, errHash := helper.HashPassword(input.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}
	userInputGorm := CoreToModel(input)
	userInputGorm.Password = hashedPassword
	tx := repo.db.Create(&userInputGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

// SelectById implements user.UserDataInterface
func (repo *userQuery) SelectById(id int) (user.Core, error) {
	var userData User
	tx := repo.db.First(&userData, id)
	if tx.Error != nil {
		return user.Core{}, tx.Error
	}

	userCore := ModelToCore(userData)

	return userCore, nil
}

// Update implements user.UserDataInterface
func (repo *userQuery) Update(id int, input user.Core) error {
	userInputGorm := CoreToModel(input)
	//userInput := models.User{}
	errUpdates := repo.db.Model(&User{}).Where("id = ?", id).Updates(UpdateUser(userInputGorm))

	if errUpdates.Error != nil {
		return errUpdates.Error
	}
	if errUpdates.RowsAffected == 0 {
		return errors.New("no rows affected, update failed")
	}
	return nil
}

// Delete implements user.UserDataInterface
func (repo *userQuery) Delete(id int) error {
	var userData User
	errDelete := repo.db.Delete(&userData, id)
	if errDelete.Error != nil {
		return errDelete.Error
	}
	return nil
}
