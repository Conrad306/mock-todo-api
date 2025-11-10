package handler

import (
	"crypto/rand"
	"net/http"
	"strconv"

	"github.com/Conrad306/mock-todo-api/models"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)


type Handler struct { 
	DB *gorm.DB
}
func (h Handler) Handle(e *echo.Echo) {
  e.Use(middleware.RemoveTrailingSlash())
  e.Static("/js", "web/js") 
  e.Static("/css", "web/css")
  e.GET("/", h.Redirect)
	
  e.GET("/:roomId", func(c echo.Context) error {
    roomId := c.Param("roomId")
    if roomId == "js" || roomId == "css" || roomId == "api" {
      return echo.ErrNotFound
    }
    return c.File("web/index.html")
  })
  
	api := e.Group("/api/:roomId")
  api.GET("/todos", h.ListTodos)
  api.POST("/todos", h.CreateTodo)
  api.PUT("/todos/:id", h.UpdateTodo)
  api.DELETE("/todos/:id", h.DeleteTodo)
}

func (h Handler) Redirect(c echo.Context) error {
	return c.Redirect(302, "/" + rand.Text())
}


func (h Handler) ListTodos(c echo.Context) error {

	roomId := c.Param("roomId") 

	var todos []models.TodoCard 
	result := h.DB.Where("room_id = ?", roomId).Find(&todos)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}
	return c.JSON(http.StatusOK, todos)
}

func (h Handler) CreateTodo(c echo.Context) error {
	title := c.FormValue("title")
	roomId := c.Param("roomId") 

	var completed = false 

	var todo = models.TodoCard {
		Title: title, 
		Completed: completed,
		RoomId: roomId,
	}

	result := h.DB.Create(&todo)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}

	return c.JSON(http.StatusOK, todo)
}

func (h Handler) UpdateTodo(c echo.Context) error {
	
	idParam := c.Param("id")
	titleFormValue := c.FormValue("title") 
	completedFormValue := c.FormValue("completed") 
	roomId := c.Param("roomId") 



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
		RoomId: roomId,
	}

	result := h.DB.Where("id = ?", id).Updates(&card)


	if result.Error != nil { 
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}


	return c.JSON(http.StatusOK, card)
}

func (h Handler) DeleteTodo(c echo.Context) error {
	roomId := c.Param("roomId") 
	idParam := c.Param("id")

	id, err := strconv.ParseInt(idParam, 10, 32)

	if err != nil { 
		return c.JSON(http.StatusInternalServerError, err.Error())
	}


	card := models.TodoCard{
		RoomId: roomId,
	}
	
	result := h.DB.Where("id = ?", id).Delete(&card)


	if result.Error != nil { 
		return c.JSON(http.StatusInternalServerError, result.Error.Error())
	}


	return c.JSON(http.StatusNoContent, "OK")
}

