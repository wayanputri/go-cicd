package handler

import (
	"be17/main/feature/project"
	"be17/main/feature/task"
	tasks "be17/main/feature/task"

	//"be17/main/feature/task/data"
	"time"
)


type ProjectResponse struct {
	Id           uint
	UserID       int
	Name_Project string
	CreatedAt     time.Time
	Task []tasks.Core
}

func CoreToResponse(input project.Core) ProjectResponse {
	return ProjectResponse{
		Id:               input.Id,
		UserID:           input.UserID,
		Name_Project:     input.Name_Project,
		CreatedAt: 		  input.CreateAt,
	}
}

func ProjectWithTask(input project.Core) project.Core {
	var tasks []task.CoreName
	for _, t := range input.Task {
		tasks = append(tasks, Task(t))
	}

	return project.Core{
		Id:           input.Id,
		Name_Project: input.Name_Project,
		UserID:       input.UserID,
		CreateAt:     input.CreateAt,
		UpdateAt:     time.Time{},
		DeleteAt:     time.Time{},
		Task:         tasks,
	}
}

func Task(input task.CoreName) task.CoreName{
	return task.CoreName{
		Name_task: input.Name_task,
	}
}