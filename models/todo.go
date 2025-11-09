package models

import (
	"gorm.io/gorm"
)

var DbConnection *gorm.DB


type TodoCard struct {
	gorm.Model
	Title       string 		`json:"title"`
	Completed   bool   		`json:"completed"`
}