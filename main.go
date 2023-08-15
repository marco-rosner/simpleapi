package main

import (
	"github.com/marco-rosner/simpleapi/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("html/**/*.html")

	router.Static("/static", "assets")
	router.StaticFile("/favicon.ico", "assets/favicon.ico")

	router.GET("/", handlers.RenderIndexHTML)
	router.GET("/user", handlers.RenderUserHTML)

	json := router.Group("/json")
	{
		json.GET("/todos", handlers.GetTodos)        // get all TODO's
		json.GET("/todos/:id", handlers.GetTodoByID) // get TODO by ID
		json.POST("/todos", handlers.PostTodo)       // post new todo
	}

	form := router.Group("/form")
	{
		form.GET("/client", handlers.AddClient) // using query string
	}

	uri := router.Group("/uri")
	{
		uri.GET("/:name/:id", handlers.AddPerson) // /:name/:uuid
	}

	router.Run("localhost:8080")
}
