package main

import (
	"diary_api/database"
	"diary_api/model"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Hello World!")

	loadEnv()
	loadDatabase()
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
