package handler

import (
	"be17/main/app/middlewares"
	"be17/main/feature/task"
	"be17/main/helper"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	TaskService task.TaskServiceInterface
}

func New(service task.TaskServiceInterface) *TaskHandler {
	return &TaskHandler{
		TaskService: service,
	}
}

func (handler *TaskHandler) Create(c echo.Context) error {
	id := middlewares.ExtracTokenUserId(c)

	taskInput := TaskRequest{}
	errBind := c.Bind(&taskInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	taskCore := RequestToCore(taskInput)

	err := handler.TaskService.Create(taskCore, id)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
}

func (handler *TaskHandler) Delete(c echo.Context) error{
	id := middlewares.ExtracTokenUserId(c)
	idProject := c.Param("taskid")
	idConv, errConv := strconv.Atoi(idProject)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}
	err := handler.TaskService.Delete(idConv,id)
	if err != nil {
		if strings.Contains(err.Error(), "delete failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete data"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error delete server data"))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}

func (handler *TaskHandler) Update(c echo.Context) error{
	id := middlewares.ExtracTokenUserId(c)

	idProject := c.Param("taskid")
	idConv, errConv := strconv.Atoi(idProject)
	if errConv != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	taskInput := TaskRequest{}
	errBind := c.Bind(&taskInput )
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	taskCore := RequestToCore(taskInput )

	err := handler.TaskService.Update(idConv,id,taskCore)
	if err != nil {
		if strings.Contains(err.Error(), "update failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error update data"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("gagal update data, server errorr"))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}