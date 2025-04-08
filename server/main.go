package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Title    string  `json:"title"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
	AisleNo  string  `json:"aisleNo"`
	Upc      string  `json:"upc"`
	Category string  `json:"category"`
}

func getItems(c *gin.Context) {
	items := []Product{
		{Title: "Test", Price: 1.99, Quantity: 50, AisleNo: "A23", Upc: "1234567890", Category: "Clothing"},
	}

	c.JSON(http.StatusOK, items)
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:5173", "https://pwa.nathan-strobl.org"},
		AllowMethods:  []string{"GET"},
		AllowHeaders:  []string{"Content-Type", "killSwitch"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	routerGroup := router.Group("/")
	{
		routerGroup.GET("", getItems)
	}

	router.Run()
}
