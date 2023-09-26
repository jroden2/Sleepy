package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jroden2/Sleepy/Echo/pkg/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Create our echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	controllers.SetupRoutes(e)

	PORT := os.Getenv("ServerPort")
	if PORT == "" {
		PORT = "8080"
	}
	hostPort := fmt.Sprintf(":%s", PORT)
	e.Logger.Fatal(e.Start(hostPort))
}
