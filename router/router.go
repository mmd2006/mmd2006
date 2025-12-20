package router

import (
	"ToDoAPP/controller"
	"ToDoAPP/middleware"
	"net/http"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "ToDoAPP API is running"})
	})
	e.GET("/admin/tasks", controller.GetTasks, middleware.JWTMiddleware, middleware.RequireRole("admin"))
	e.GET("/admin/users", controller.GetUsers, middleware.JWTMiddleware, middleware.RequireRole("admin"))

	e.POST("/tasks", controller.CreateTask, middleware.JWTMiddleware)
	e.GET("/tasks", controller.GetTasks, middleware.JWTMiddleware)
	e.GET("/tasks/:id", controller.GetTaskByID, middleware.JWTMiddleware)
	e.PUT("/tasks/:id", controller.UpdateTask, middleware.JWTMiddleware)
	e.DELETE("/tasks/:id", controller.DeleteTask, middleware.JWTMiddleware)

}

func AuthRoutes(e *echo.Echo) {
	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)
}
