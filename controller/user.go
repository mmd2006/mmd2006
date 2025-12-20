package controller

import (
	"ToDoAPP/config"
	"ToDoAPP/model"
	"ToDoAPP/validation"
	"context"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c echo.Context) error {
	var input validation.SignupInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
		Role:     model.RoleUser,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := config.UserCollection.CountDocuments(ctx, bson.M{"username": user.Username})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "database error"})
	}

	if count > 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "username already exists"})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to hash password"})
	}
	user.Password = string(hashedPassword)

	_, err = config.UserCollection.InsertOne(ctx, user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "database error"})
	}
	return c.JSON(http.StatusCreated, echo.Map{"message": "user created"})
}

func Login(c echo.Context) error {
	var input validation.SignupInput
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	if err := input.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid input"})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := config.UserCollection.FindOne(ctx, bson.M{"username": input.Username}).Decode(&user)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credentials"})
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"user_id":  user.ID.Hex(),
		"role":     user.Role,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtSecret := os.Getenv("JWT_SECRET")
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to sign JWT"})
	}

	return c.JSON(http.StatusOK, echo.Map{"token": signedToken})
}

func GetUsers(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.UserCollection.Find(ctx, bson.M{})
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": "database error"})
	}

	var users []model.User
	if err = cursor.All(ctx, &users); err != nil {
		return e.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	response := make([]map[string]interface{}, len(users))
	for i, u := range users {
		response[i] = map[string]interface{}{
			"id":       u.ID.Hex(),
			"username": u.Username,
			"role":     u.Role,
		}
	}

	return e.JSON(http.StatusOK, echo.Map{"users": response})
}
