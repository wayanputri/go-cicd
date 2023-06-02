package data

import (
	"be17/main/feature/task"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID int
	ProjectID int
	Name_task   string `validate:"required"`
	Status_task string
	Data        string
}

func CoreToModel(dataCore task.Core)Task{
	return  Task{
		Model:        gorm.Model{},
		UserID:       dataCore.UserID,
		ProjectID: dataCore.ProjectID,
		Name_task: dataCore.Name_task,
		Status_task: dataCore.Status_task,
		Data: dataCore.Data,
	}
}

func ModelToModel(dataCore Task)Task{
	return  Task{
		Model:        gorm.Model{},
		UserID:       dataCore.UserID,
		ProjectID: dataCore.ProjectID,
		Name_task: dataCore.Name_task,
		Status_task: dataCore.Status_task,
		Data: dataCore.Data,
	}
}
