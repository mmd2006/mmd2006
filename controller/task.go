package controller

import (
	"ToDoAPP/config"
	"ToDoAPP/model"
	"ToDoAPP/validation"
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTask(c echo.Context) error {
	var input validation.TaskInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user := c.Get("user")
	claims := user.(jwt.MapClaims)

	userIDstr := claims["user_id"].(string)
	userID, err := primitive.ObjectIDFromHex(userIDstr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user ID"})
	}

	task := model.Task{
		ID:          primitive.NewObjectID(),
		UserID:      userID,
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = config.TaskCollection.InsertOne(ctx, task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "database error"})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"id":          task.ID.Hex(),
		"user_id":     task.UserID.Hex(),
		"title":       task.Title,
		"description": task.Description,
		"completed":   task.Completed,
	})
}

func GetTasks(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := c.Get("user").(jwt.MapClaims)
	userIDHex := user["user_id"].(string)

	role, ok := user["role"].(string)
	if !ok {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid role in token"})
	}

	var filter bson.M
	if role == model.RoleAdmin {
		filter = bson.M{}
	} else {
		userID, err := primitive.ObjectIDFromHex(userIDHex)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user ID"})
		}
		filter = bson.M{"user_id": userID}
	}

	cursor, err := config.TaskCollection.Find(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	var tasks []model.Task
	if err = cursor.All(ctx, &tasks); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	response := make([]map[string]interface{}, len(tasks))
	for i, t := range tasks {
		response[i] = map[string]interface{}{
			"id":          t.ID.Hex(),
			"user_id":     t.UserID.Hex(),
			"title":       t.Title,
			"description": t.Description,
			"completed":   t.Completed,
		}
	}

	return c.JSON(http.StatusOK, response)
}

func GetTaskByID(c echo.Context) error {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid task id"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	user := c.Get("user").(jwt.MapClaims)
	userIDHex := user["user_id"].(string)
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid user id"})
	}

	filter := bson.M{"_id": objectID, "user_id": userID}

	var task model.Task
	err = config.TaskCollection.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "task not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"id":          task.ID.Hex(),
		"user_id":     task.UserID.Hex(),
		"title":       task.Title,
		"description": task.Description,
		"completed":   task.Completed,
	})
}

func UpdateTask(c echo.Context) error {
	idparam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idparam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID format"})
	}

	var input validation.TaskInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}
	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	user := c.Get("user").(jwt.MapClaims)
	userIDHex := user["user_id"].(string)
	userID, _ := primitive.ObjectIDFromHex(userIDHex)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID, "user_id": userID}
	update := bson.M{
		"$set": bson.M{
			"title":       input.Title,
			"description": input.Description,
		},
	}

	result, err := config.TaskCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	if result.MatchedCount == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "task not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "task updated"})
}

func DeleteTask(c echo.Context) error {
	idparam := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(idparam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid ID format"})
	}

	user := c.Get("user").(jwt.MapClaims)
	userIDHex := user["user_id"].(string)
	userID, _ := primitive.ObjectIDFromHex(userIDHex)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": objectID, "user_id": userID}

	result, err := config.TaskCollection.DeleteOne(ctx, filter)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	if result.DeletedCount == 0 {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "task not found"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "task deleted"})

}
