package service

import (
	"be17/main/feature/project"

	"github.com/go-playground/validator/v10"
)

type ProjectService struct {
	projectData project.ProjectDataInterface
	validate    *validator.Validate
}

// Delete implements project.ProjectServiceInterface
func (service *ProjectService) Delete(idUser int, idProject int) error {
	err := service.projectData.Delete(idUser,idProject)
	return err
}

// Update implements project.ProjectServiceInterface
func (service *ProjectService) Update(idUser int, idProject int, input project.Core) error {
	err := service.projectData.UpdateData(idUser, idProject, input)
	return err
}

// GetByIdProject implements project.ProjectServiceInterface
func (service *ProjectService) GetById(idUser int, idProject int) (project.Core, error) {
	data, err := service.projectData.SelectById(idUser, idProject)
	return data, err
}

// GetById implements project.ProjectServiceInterface
func (service *ProjectService) GetAll(id int) ([]project.Core, error) {
	data, err := service.projectData.SelectAll(id)
	return data, err
}

// Create implements project.ProjectServiceInterface
func (service *ProjectService) Create(input project.Core, id int) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}

	// Menyimpan data user ke database
	errInsert := service.projectData.Insert(input, id)
	return errInsert
}

func New(repo project.ProjectDataInterface) project.ProjectServiceInterface {
	return &ProjectService{
		projectData: repo,
		validate:    validator.New(),
	}
}
