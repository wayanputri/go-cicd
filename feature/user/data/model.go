package data

import (
	"be17/main/feature/user"

	"gorm.io/gorm"
)
type User struct{
	gorm.Model
	Name     string
	Phone    string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string

}

func CoreToModel(dataCore user.Core)User{
	return  User{
			Name:     dataCore.Name,
			Phone:    dataCore.Phone,
			Email:    dataCore.Email,
			Password: dataCore.Password,
		}
}

func ModelToCore(dataModel User) user.Core{
	return  user.Core{
			Id: dataModel.ID,
			Name:     dataModel.Name,
			Phone:    dataModel.Phone,
			Email:    dataModel.Email,
			Password: dataModel.Password,
			CreatedAt: dataModel.CreatedAt,
			UpdatedAt: dataModel.UpdatedAt,
		}
}

func UpdateUser(input User)User{
	return User{
		Name: input.Name,
		Phone: input.Phone,
		Email: input.Email,
		Password: input.Password,
	}
}