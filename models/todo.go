package models

import (
	"gorm.io/gorm"
)

type TodoCard struct {
	gorm.Model
	RoomId			string		`json:"roomId"`
	Title       string 		`json:"title"`
	Completed   bool   		`json:"completed"`
}
