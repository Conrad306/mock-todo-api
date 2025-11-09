package main

import (
	"log"
	"os"

	"github.com/Conrad306/mock-todo-api/internal/handler"
	"github.com/Conrad306/mock-todo-api/models"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func main() {
	e := echo.New()

	err := godotenv.Load()

	if err != nil { 
		log.Fatal("Failed to load .ENV file: ", err)
	}

	dsn := os.Getenv("CONNECTION_STRING")


	db, err := gorm.Open(postgres.Open(dsn))

	db.AutoMigrate(&models.TodoCard{})

	if err != nil { 
		log.Fatal("Failed to load DB: ", err)
	}

	models.DbConnection = db
	handler.Handler(e)

	e.Logger.Fatal(e.Start(":8080"))
}