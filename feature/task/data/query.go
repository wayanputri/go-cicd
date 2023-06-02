package data

import (
	"be17/main/feature/task"
	"errors"

	"gorm.io/gorm"
)

type taskQuery struct {
	db *gorm.DB
}

// Update implements task.TaskDataInterface
func (repo *taskQuery) Update(id int, UserId int, input task.Core) error {
	taskInputGorm := CoreToModel(input)
	//userInput := models.User{}
	errUpdates := repo.db.Model(&Task{}).Where("id = ? AND user_id =?", id,UserId).Updates(ModelToModel(taskInputGorm))

	if errUpdates.Error != nil {
		return errUpdates.Error
	}
	if errUpdates.RowsAffected == 0 {
		return errors.New("no rows affected, update failed")
	}
	return nil
}

// Delete implements task.TaskDataInterface
func (repo *taskQuery) Delete(id int, UserId int) error {
	var taskData Task
	errDelete := repo.db.Delete(&taskData, "id=? AND user_id=?", id, UserId)
	if errDelete.Error != nil {
		return errDelete.Error
	}
	return nil
}

// Insert implements task.TaskDataInterface
func (repo *taskQuery) Insert(input task.Core, userId int) error {
	taskInputGorm := CoreToModel(input)
	taskInputGorm.UserID = userId
	tx := repo.db.Create(&taskInputGorm)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

func New(db *gorm.DB) task.TaskDataInterface {
	return &taskQuery{
		db: db,
	}
}
