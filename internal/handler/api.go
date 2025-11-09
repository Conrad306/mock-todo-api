package handler

import (
	"net/http"
	"strconv"

	"github.com/Conrad306/mock-todo-api/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Handler(e *echo.Echo) {

	e.Use(middleware.RemoveTrailingSlash())

	todo := e.Group("todo")


	todo.GET("/list", ListTodos)
	todo.POST("/create", CreateTodo)
	todo.PATCH("/update/:id", UpdateTodo)
	todo.DELETE("/delete/:id", DeleteTodo)	
	
}

func ListTodos(c echo.Context) error {
	var todos []models.TodoCard 
	result := models.DbConnection.Find(&todos)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

func CreateTodo(c echo.Context) error {
	title := c.FormValue("title")
	var completed = false 

	var todo = models.TodoCard {
		Title: title, 
		Completed: completed,
	}

	result := models.DbConnection.Create(&todo)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c echo.Context) error {
	
	idParam := c.Param("id")
	titleFormValue := c.FormValue("title") 
	completedFormValue := c.FormValue("completed") 


	id, err := strconv.ParseInt(idParam, 10, 32)

	if err != nil { 
		return c.JSON(http.StatusInternalServerError, err.Error())
	}


	completed, err := strconv.ParseBool(completedFormValue)

	
	if err != nil { 
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	card := models.TodoCard {
		Title: titleFormValue,
		Completed: completed,
	}

	result := models.DbConnection.Where("id = ?", id).Updates(&card)


	if result.Error != nil { 
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}


	return c.JSON(http.StatusOK, card)
}

func DeleteTodo(c echo.Context) error {
	
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 32)

	if err != nil { 
		return c.JSON(http.StatusInternalServerError, err.Error())
	}


	card := models.TodoCard{}
	
	result := models.DbConnection.Where("id = ?", id).Delete(&card)


	if result.Error != nil { 
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}


	return c.JSON(http.StatusNoContent, "OK")
}

