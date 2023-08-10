package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":       "Go rendering HTML",
		"description": "by Marco Rosner",
	})
}
