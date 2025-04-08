package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func getItems(c *gin.Context) {

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
}
