package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RenderIndexHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{
		"title": "Go rendering HTML",
		"page":  "index.html",
	})
}

func RenderUserHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "user/user.html", gin.H{
		"title": "Go rendering HTML",
		"page":  "user.html",
	})
}
