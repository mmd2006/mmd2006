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

	adminGroup := e.Group("/admin", middleware.JWTMiddleware)
	adminGroup.GET("/admin/tasks", controller.GetTasks, middleware.RequireRole("admin"))
	adminGroup.GET("/admin/users", controller.GetUsers, middleware.RequireRole("admin"))

	userGroup := e.Group("/tasks", middleware.JWTMiddleware)
	userGroup.POST("/tasks", controller.CreateTask)
	userGroup.GET("/tasks", controller.GetTasks)
	userGroup.GET("/tasks/:id", controller.GetTaskByID)
	userGroup.PUT("/tasks/:id", controller.UpdateTask)
	userGroup.DELETE("/tasks/:id", controller.DeleteTask)

}

func AuthRoutes(e *echo.Echo) {
	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)
}
