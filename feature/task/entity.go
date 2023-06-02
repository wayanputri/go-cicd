package task

import "time"

type Core struct {
	Id         uint
	UserID int
	ProjectID int
	Name_task string `validate:"required"`
	Status_task string
	Data string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type CoreName struct {
	Name_task string `validate:"required"`
}

type TaskDataInterface interface{
	Insert(input Core,userId int) error
	Update(id int,UserId int,input Core) error
	Delete(id int,UserId int) error
}

type TaskServiceInterface interface{
	Create(input Core,userId int) error
	Update(id int,UserId int,input Core) error
	Delete(id int,UserId int) error
}