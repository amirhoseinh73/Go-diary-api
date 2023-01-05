package main

import (
	"diary_api/controller"
	"diary_api/database"
	"diary_api/model"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World!")

	loadEnv()
	loadDatabase()
	serverApplication()
}

func loadEnv() {
	err := godotenv.Load(".env.local")

	if err != nil {
		fmt.Println("Error loading env file")
		os.Exit(1)
	}

	fmt.Println("env loaded successfully.")
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.Entry{})
	database.Database.AutoMigrate(&model.User{})
}

func serverApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	PORT := fmt.Sprintf(":%s", os.Getenv("APP_PORT"))

	router.Run(PORT)
	fmt.Println("Server Running on port " + PORT)
}
