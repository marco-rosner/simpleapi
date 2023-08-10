package main

import (
	"html/template"

	"github.com/simpleapi/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFile("/favicon.ico", "assets/favicon.ico")
	html := template.Must(template.ParseGlob("html/*"))
	router.SetHTMLTemplate(html)

	router.GET("/", handlers.RenderHTML)

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
