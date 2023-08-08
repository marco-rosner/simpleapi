package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// json routes
	router.GET("/todos", getTodos)        // get all TODO's
	router.GET("/todos/:id", getTodoByID) // get TODO by ID
	router.POST("/todos", postTodo)       // post new todo

	// form routes
	router.GET("/client", addClient) // using query string

	// uri routes
	router.GET("/:name/:id", addPerson) // /:name/:uuid

	router.Run("localhost:8080")
}
