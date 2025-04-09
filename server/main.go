package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ProductCategory struct {
	Title    string    `json:"title"`
	Products []Product `json:"products"`
}

type Product struct {
	Title    string  `json:"title"`
	Price    float32 `json:"price"`
	Quantity int     `json:"quantity"`
	AisleNo  string  `json:"aisleNo"`
	Upc      string  `json:"upc"`
}

func getAllProductCategories() []ProductCategory {
	return []ProductCategory{
		{Title: "Test", Products: []Product{
			{Title: "Test", Price: 1.99, Quantity: 50, AisleNo: "A23", Upc: "1234567890"},
		}},
		{Title: "Test2", Products: []Product{
			{Title: "iPhone 14 Pro Max (128GB)", Price: 1199.99, Quantity: 10, AisleNo: "A23", Upc: "000000000000"},
		}},
	}
}

func getProductCategory(c *gin.Context) {
	productCategories := getAllProductCategories()
	var selectedCategory ProductCategory
	categoryFound := false

	categoryUrlParam := c.Query("category")

	for _, element := range productCategories {
		if element.Title == categoryUrlParam {
			selectedCategory = element
			categoryFound = true
		}
	}

	if categoryFound {
		c.JSON(http.StatusOK, selectedCategory)
	} else {
		c.JSON(http.StatusNotFound, "Category was not specified or the requested product category does not exist.")
	}
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"http://localhost:4173", "http://localhost:5173", "https://pwa.nathan-strobl.org"},
		AllowMethods:  []string{"GET"},
		AllowHeaders:  []string{"Content-Type", "killSwitch"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	routerGroup := router.Group("/")
	{
		routerGroup.GET("", getProductCategory)
	}

	router.Run()
}
