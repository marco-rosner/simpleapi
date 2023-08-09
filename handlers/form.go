package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Client struct {
	Name         string    `form:"name"`
	Address      string    `form:"address"`
	Birthday     time.Time `form:"birthday" binding:"required"` // Format: 2022-10-03T00:00:00Z
	PaymentDate  time.Time `form:"payment_date" binding:"required"`
	DeliveryDate time.Time `form:"delivery_date" binding:"required,gtefield=PaymentDate"`
}

func AddClient(c *gin.Context) {
	var client Client

	if err := c.ShouldBind(&client); err != nil {
		log.Println(client.Name)
		log.Println(client.Address)
		log.Println(client.Birthday)
		log.Println(client.DeliveryDate)
		log.Println(client.PaymentDate)

		c.JSON(http.StatusBadRequest, gin.H{"ValidationError": err.Error()})
	}

	jsonClient, err := json.Marshal(client)

	if err != nil {
		log.Println(err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
	}

	c.IndentedJSON(http.StatusOK, string(jsonClient))
}
