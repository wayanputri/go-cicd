package handler

import (
	"be17/main/app/middlewares"
	"be17/main/feature/project"
	"be17/main/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	ProjectService project.ProjectServiceInterface
}

func New(service project.ProjectServiceInterface) *ProjectHandler {
	return &ProjectHandler{
		ProjectService: service,
	}
}

func (handler *ProjectHandler) Create(c echo.Context) error {
	id := middlewares.ExtracTokenUserId(c)

	userInput := ProjectRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	userCore := RequestToCore(userInput)

	err := handler.ProjectService.Create(userCore,id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}

func (handler *ProjectHandler) GetAll(c echo.Context) error{

	id := middlewares.ExtracTokenUserId(c)
	result, err := handler.ProjectService.GetAll(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	var projectResponseAll []ProjectResponse
	for _,value := range result{
		projectResponse := CoreToResponse(value)
		projectResponseAll = append(projectResponseAll, projectResponse)
	}
	

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data",projectResponseAll ))
}

func (handler *ProjectHandler) GetById(c echo.Context) error{

	idUser := middlewares.ExtracTokenUserId(c)
	idProject := c.Param("projectid")
	idConv, errConv := strconv.Atoi(idProject)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	result, err := handler.ProjectService.GetById(idUser,idConv)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
		projectResponse := ProjectWithTask(result)

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data",projectResponse))
}

func (handler *ProjectHandler) Update(c echo.Context) error{
	idUser := middlewares.ExtracTokenUserId(c)
	idProject := c.Param("projectid")
	idConv, errConv := strconv.Atoi(idProject)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	projectInput := ProjectRequest{}
	errBind := c.Bind(&projectInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	projectCore := RequestToCore(projectInput)

	err := handler.ProjectService.Update(idUser,idConv,projectCore)
	if err != nil {
		if strings.Contains(err.Error(), "update failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error update data"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("gagal update data, server errorr"))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

func (handler *ProjectHandler) Delete(c echo.Context) error{
	id := middlewares.ExtracTokenUserId(c)

	idProject := c.Param("projectid")
	idConv, errConv := strconv.Atoi(idProject)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	
	err := handler.ProjectService.Delete(id,idConv)
	if err != nil {
		if strings.Contains(err.Error(), "delete failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete data"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error delete server data"))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}


