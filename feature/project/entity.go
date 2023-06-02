package project

import (
	"be17/main/feature/task"
	"time"
)

type Core struct {
	Id           uint
	Name_Project string
	UserID int
	CreateAt     time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	Task []task.CoreName
}

type ProjectDataInterface interface{
	Insert(input Core,id int) error
	SelectAll(id int) ([]Core,error)
	SelectById(idUser int, idProject int) (Core,error)
	UpdateData(idUser int, idProject int,input Core) error
	Delete(idUser int, idProject int) error
}

type ProjectServiceInterface interface{
	Create(input Core,id int) error
	GetAll(id int) ([]Core,error)
	GetById(idUser int, idProject int) (Core,error)
	Update(idUser int, idProject int,input Core) error
	Delete(idUser int, idProject int) error
}