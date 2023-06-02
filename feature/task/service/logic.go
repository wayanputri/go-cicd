package service

import (
	"be17/main/feature/task"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type taskService struct {
	taskData task.TaskDataInterface
	validate *validator.Validate
}

// Update implements task.TaskServiceInterface
func (service *taskService) Update(id int, UserId int, input task.Core) error {
	err := service.taskData.Update(id,UserId, input)
	if err != nil {
		return fmt.Errorf("failed to update user with ID %d: %w", id, err)
	}
	return nil
}

// Delete implements task.TaskServiceInterface
func (service *taskService) Delete(id int, UserId int) error {
	err := service.taskData.Delete(id, UserId)
	return err
}

// Create implements task.TaskServiceInterface
func (service *taskService) Create(input task.Core, userId int) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	// Menyimpan data user ke database
	errInsert := service.taskData.Insert(input, userId)
	return errInsert
}

func New(repo task.TaskDataInterface) task.TaskServiceInterface {
	return &taskService{
		taskData: repo,
		validate: validator.New(),
	}
}
