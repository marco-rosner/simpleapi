package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Client struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"10-03-2023" time_utc:"1"`
}

func AddClient(c *gin.Context) {
	var client Client

	if c.ShouldBind(&client) == nil {
		log.Println(client.Name)
		log.Println(client.Address)
		log.Println(client.Birthday)
	}

	jsonClient, err := json.Marshal(client)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}

	c.IndentedJSON(http.StatusOK, string(jsonClient))
}
