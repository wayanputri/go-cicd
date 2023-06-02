package handler

type TaskRequest struct {
	Name_task   string `json:"name_task" form:"name_task"`
	ProjectID   int    `json:"project_id" form:"project_id"`
	Status_task string `json:"status_task" form:"status_task"`
	Data        string `json:"data" form:"data"`
}
