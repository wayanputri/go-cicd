package data

import (
	"be17/main/feature/project"
	"errors"

	"gorm.io/gorm"
)

type projectQuery struct {
	db *gorm.DB
}

// Delete implements project.ProjectDataInterface
func (repo *projectQuery) Delete(idUser int, idProject int) error {
	var projectData Project
	errDelete := repo.db.Delete(&projectData, "user_id = ? AND id = ?", idUser, idProject)
	if errDelete.Error != nil {
		return errDelete.Error
	}
	return nil
}

// Update implements project.ProjectDataInterface
func (repo *projectQuery) UpdateData(idUser int, idProject int, input project.Core) error {
	projectInputGorm := CoreToModel(input)
	//userInput := models.User{}
	errUpdates := repo.db.Model(&Project{}).Where("id = ? AND user_id=?", idProject, idUser).Updates(UpdateModelToModel(projectInputGorm))

	if errUpdates.Error != nil {
		return errUpdates.Error
	}
	if errUpdates.RowsAffected == 0 {
		return errors.New("no rows affected, update failed")
	}
	return nil
}

// SelectByIdProject implements project.ProjectDataInterface
func (repo *projectQuery) SelectById(idUser int, idProject int) (project.Core, error) {

	var projectData Project
	result := repo.db.
		Preload("Task").
		Where("user_id = ? AND id = ?", idUser, idProject).
		First(&projectData)
	if result.Error != nil {
		return project.Core{}, result.Error
	}

	projectWithTasks := ProjectWithTask(projectData)
	return projectWithTasks, nil
}

// SelectById implements project.ProjectDataInterface
func (repo *projectQuery) SelectAll(id int) ([]project.Core, error) {
	var projectDataAll []Project
	
	tx := repo.db.Find(&projectDataAll, "user_id = ?", id)
	if tx.Error != nil {
		return []project.Core{}, tx.Error
	}

	var projectAll []project.Core

	for _, value := range projectDataAll {
		projectCore := ModelToCore(value)
		projectAll = append(projectAll, projectCore)
	}

	return projectAll, nil

}

// Insert implements project.ProjectDataInterface
func (repo *projectQuery) Insert(input project.Core, id int) error {

	projectInputGorm := CoreToModel(input)
	projectInputGorm.UserID = id
	tx := repo.db.Create(&projectInputGorm)

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}
	return nil
}

func New(db *gorm.DB) project.ProjectDataInterface {
	return &projectQuery{
		db: db,
	}
}
