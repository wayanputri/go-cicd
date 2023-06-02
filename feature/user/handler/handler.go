package handler

import (
	"be17/main/app/middlewares"
	"be17/main/feature/user"
	"be17/main/helper"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct{
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserHandler{
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) CreateUser(c echo.Context) error{

	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil{
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}

	userCore := RequestToCore(userInput)	

	err := handler.userService.Create(userCore)
	if err != nil{
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse(err.Error()))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error insert data"+err.Error()))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success insert data"))
	}

func (handler *UserHandler) Login(c echo.Context) error{
	loginInput := AuthRequest{}
	errBind := c.Bind(&loginInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	dataLogin, token,err :=handler.userService.Login(loginInput.Email,loginInput.Password)
	if err != nil{
		if strings.Contains(err.Error(),"login failed"){
			return c.JSON(http.StatusBadRequest,helper.FailedResponse(err.Error()))
		}else{
			return c.JSON(http.StatusInternalServerError,helper.FailedResponse("error login, internal server error"))
		}
	}
	return c.JSON(http.StatusOK,helper.SuccessWithDataResponse("login succes", map[string]any{
		"token": token,
		"email": dataLogin.Email,
		"id": dataLogin.Id,
	}))
}

func (handler *UserHandler) GetById(c echo.Context) error{

	id := middlewares.ExtracTokenUserId(c)
	result, err := handler.userService.GetById(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error read data"))
	}

	usersResponse := CoreToResponse(result)

	// response ketika berhasil
	return c.JSON(http.StatusOK, helper.SuccessWithDataResponse("success read data",usersResponse ))
}

func (handler *UserHandler) Update(c echo.Context) error{
	id := middlewares.ExtracTokenUserId(c)

	userInput := UserRequest{}
	errBind := c.Bind(&userInput)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, helper.FailedResponse("error bind data"))
	}
	userCore := RequestToCore(userInput)

	err := handler.userService.Update(id,userCore)
	if err != nil {
		if strings.Contains(err.Error(), "update failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error update data"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("gagal update data, server errorr"))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success update data"))
}

func (handler *UserHandler) Delete(c echo.Context) error{
	id := middlewares.ExtracTokenUserId(c)
	err := handler.userService.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "delete failed") {
			return c.JSON(http.StatusBadRequest, helper.FailedResponse("error delete data"))
		} else {
			return c.JSON(http.StatusInternalServerError, helper.FailedResponse("error delete server data"))
		}
	}
	return c.JSON(http.StatusOK, helper.SuccessResponse("success delete data"))
}

