package main

import (
	"ToDoAPP/config"
	"ToDoAPP/router"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, using system env")
	}

	config.ConnectMongo()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.AuthRoutes(e)
	router.InitRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	log.Fatal(e.Start(":" + port))

}
