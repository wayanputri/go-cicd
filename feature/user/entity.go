package user

import (
	"time"
)

type Core struct {
	Id        uint
	Name      string `validate:"required"`
	Phone     string
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserDataInterface interface{
	Insert(input Core) error
	Login(email string, password string) (Core, string, error)
	SelectById(id int) (Core,error)
	Update(id int,input Core) error
	Delete(id int) error
}

type UserServiceInterface interface{
	Create(input Core) error
	Login(email string, password string) (Core, string, error)
	GetById(id int) (Core,error)
	Update(id int,input Core) error
	Delete(id int) error
}