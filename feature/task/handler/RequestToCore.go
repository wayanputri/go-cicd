package handler

import "be17/main/feature/task"

func RequestToCore(input TaskRequest) task.Core{
	return task.Core{
		ProjectID:   input.ProjectID,
		Name_task:   input.Name_task,
		Status_task: input.Status_task,
		Data:        input.Data,
	}
}