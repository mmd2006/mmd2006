package main

import (
	"ToDoAPP/config"
	"ToDoAPP/router"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"log"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println()
	}

	config.ConnectMongo()

	e := echo.New()

	router.AuthRoutes(e)
	router.InitRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	log.Fatal(e.Start(":" + port))

}
