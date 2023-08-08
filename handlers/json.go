package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todos = []Todo{
	{ID: "1", Title: "Activity 1", Description: "Description 1"},
	{ID: "2", Title: "Activity 2", Description: "Description 2"},
	{ID: "3", Title: "Activity 3", Description: "Description 3"},
}

func GetTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

func PostTodo(c *gin.Context) {
	var newTodo Todo

	if err := c.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	c.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodoByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range todos {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "TODO not found"})
}
