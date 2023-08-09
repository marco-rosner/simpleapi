package main

import (
	"github.com/simpleapi/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// json routes
	router.GET("/todos", handlers.GetTodos)        // get all TODO's
	router.GET("/todos/:id", handlers.GetTodoByID) // get TODO by ID
	router.POST("/todos", handlers.PostTodo)       // post new todo

	// form routes
	router.GET("/client", handlers.AddClient) // using query string

	// uri routes
	router.GET("/:name/:id", handlers.AddPerson) // /:name/:uuid

	router.Run("localhost:8080")
}
