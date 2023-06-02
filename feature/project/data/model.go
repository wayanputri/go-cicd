package data

import (
	"be17/main/feature/project"
	"be17/main/feature/task"
	"be17/main/feature/task/data"
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	UserID int
	Name_Project string
	Task []data.Task
}

func CoreToModel(dataCore project.Core)Project{
	return  Project{
		Model:        gorm.Model{},
		UserID:       dataCore.UserID,
		Name_Project: dataCore.Name_Project,
	}
}

func ModelToCore(dataCore Project)project.Core{
	return  project.Core{
		Id:           dataCore.ID,
		Name_Project: dataCore.Name_Project,
		UserID:       dataCore.UserID,
		CreateAt:     time.Time{},
	}
}

func UpdateModelToModel(input Project) Project{
	return Project{
		Model:        gorm.Model{},
		UserID:       0,
		Name_Project: input.Name_Project,
	} 
}

func ProjectWithTask(input Project) project.Core{
	var tasks []task.CoreName
	for _, task := range input.Task {
		tasks = append(tasks, FoundTask(task))
	}

	return project.Core{
		Id:           input.ID,
		Name_Project: input.Name_Project,
		UserID:       input.UserID,
		CreateAt:     input.CreatedAt,
		UpdateAt:     time.Time{},
		DeleteAt:     time.Time{},
		Task:         tasks,
	}
}

func FoundTask(input data.Task) task.CoreName{
	return task.CoreName{
		Name_task: input.Name_task,
	}
}