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
	adminGroup.GET("/tasks", controller.GetTasks, middleware.RequireRole("admin"))
	adminGroup.GET("/users", controller.GetUsers, middleware.RequireRole("admin"))

	userGroup := e.Group("/tasks", middleware.JWTMiddleware)
	userGroup.POST("", controller.CreateTask)
	userGroup.GET("", controller.GetTasks)
	userGroup.GET("/:id", controller.GetTaskByID)
	userGroup.PUT("/:id", controller.UpdateTask)
	userGroup.DELETE("/:id", controller.DeleteTask)

}

func AuthRoutes(e *echo.Echo) {
	e.POST("/signup", controller.Signup)
	e.POST("/login", controller.Login)
}
