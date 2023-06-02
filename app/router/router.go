package router

import (
	"be17/main/app/middlewares"
	_userData "be17/main/feature/user/data"
	_userHandler "be17/main/feature/user/handler"
	_userServis "be17/main/feature/user/service"

	_projectData "be17/main/feature/project/data"
	_projectHandler "be17/main/feature/project/handler"
	_projectServis "be17/main/feature/project/service"

	_taskData "be17/main/feature/task/data"
	_taskHandler "be17/main/feature/task/handler"
	_taskServis "be17/main/feature/task/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB,e *echo.Echo){
	userData := _userData.New(db)
	userService := _userServis.New(userData)
	userHandlerAPI := _userHandler.New(userService)
	e.POST("/users",userHandlerAPI.CreateUser)
	e.POST("/login",userHandlerAPI.Login)
	e.GET("/users",userHandlerAPI.GetById,middlewares.JWTMiddleware())
	e.PUT("/users",userHandlerAPI.Update,middlewares.JWTMiddleware())
	e.DELETE("/users",userHandlerAPI.Delete,middlewares.JWTMiddleware())

	projectData := _projectData.New(db)
	projectService := _projectServis.New(projectData)
	projectHandlerAPI := _projectHandler.New(projectService)
	e.GET("/projects",projectHandlerAPI.GetAll,middlewares.JWTMiddleware())
	e.POST("/projects",projectHandlerAPI.Create,middlewares.JWTMiddleware())
	e.GET("/projects/:projectid",projectHandlerAPI.GetById,middlewares.JWTMiddleware())
	e.PUT("/projects/:projectid",projectHandlerAPI.Update,middlewares.JWTMiddleware())
	e.DELETE("/projects/:projectid",projectHandlerAPI.Delete,middlewares.JWTMiddleware())

	taskData := _taskData.New(db)
	taskService := _taskServis.New(taskData)
	taskHandlerAPI := _taskHandler.New(taskService)
	//e.GET("/tasks",taskHandlerAPI.GetAll,middlewares.JWTMiddleware())
	e.POST("/tasks",taskHandlerAPI.Create,middlewares.JWTMiddleware())
	e.DELETE("/tasks/:taskid",taskHandlerAPI.Delete,middlewares.JWTMiddleware())
	e.PUT("/tasks/:taskid",taskHandlerAPI.Update,middlewares.JWTMiddleware())
	
}