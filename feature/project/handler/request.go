package handler

import "be17/main/feature/project"

type ProjectRequest struct {
	Name_Project string `json:"name_project" form:"name_project"`
	UserID       int    `json:"user_id" form:"user_id"`
}

func RequestToCore(dataInput ProjectRequest) project.Core {
	return project.Core{
		Name_Project:     dataInput.Name_Project,
		UserID:  dataInput.UserID,
	}
}

