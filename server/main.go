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
		{Title: "Clothing", Products: []Product{
			{Title: "Shirt", Price: 10.99, Quantity: 50, AisleNo: "A23", Upc: "1234567890"},
			{Title: "Pants", Price: 12.99, Quantity: 50, AisleNo: "A23", Upc: "3214567890"},
			{Title: "Sweater", Price: 15.99, Quantity: 50, AisleNo: "A23", Upc: "1234567897"},
			{Title: "Graphic Tee", Price: 11.99, Quantity: 50, AisleNo: "A23", Upc: "1234567899"},
			{Title: "Turtleneck Sweater", Price: 13.99, Quantity: 50, AisleNo: "A23", Upc: "1234567895"},
			{Title: "Skirt", Price: 10.99, Quantity: 50, AisleNo: "A23", Upc: "1234567891"},
			{Title: "Socks", Price: 5.99, Quantity: 50, AisleNo: "A23", Upc: "1234567892"},
			{Title: "Beanie", Price: 9.99, Quantity: 50, AisleNo: "A23", Upc: "1234567893"},
		}},
		{Title: "Home Goods", Products: []Product{
			{Title: "Sofa", Price: 1199.99, Quantity: 5, AisleNo: "B23", Upc: "000000000001"},
			{Title: "Chair", Price: 25.99, Quantity: 6, AisleNo: "B23", Upc: "000000000002"},
			{Title: "Desk", Price: 699.99, Quantity: 7, AisleNo: "B23", Upc: "000000000003"},
			{Title: "Gamer Chair", Price: 10000.99, Quantity: 8, AisleNo: "B23", Upc: "000000000004"},
			{Title: "Bookshelf", Price: 1199.99, Quantity: 10, AisleNo: "B23", Upc: "000000000005"},
			{Title: "Dresser", Price: 799.99, Quantity: 10, AisleNo: "B23", Upc: "000000000006"},
			{Title: "Mirror", Price: 299.99, Quantity: 9, AisleNo: "B23", Upc: "000000000007"},
			{Title: "Poker Mat", Price: 49.99, Quantity: 10, AisleNo: "B23", Upc: "000000000008"},
		}},
		{Title: "Electronics", Products: []Product{
			{Title: "iPhone 14 Pro Max (128GB)", Price: 1199.99, Quantity: 15, AisleNo: "C23", Upc: "100000000000"},
			{Title: "HMD Vibe (128GB)", Price: 1199.99, Quantity: 10, AisleNo: "C23", Upc: "200000000000"},
			{Title: "iPhone 13 (64GB)", Price: 899.99, Quantity: 8, AisleNo: "C23", Upc: "300000000000"},
			{Title: "Google Home", Price: 159.99, Quantity: 11, AisleNo: "C23", Upc: "400000000000"},
			{Title: "Amazon Alexa", Price: 169.99, Quantity: 10, AisleNo: "C23", Upc: "500000000000"},
			{Title: "Raspberry Pi 5", Price: 119.99, Quantity: 9, AisleNo: "C23", Upc: "600000000000"},
			{Title: "Arduino Uno", Price: 39.99, Quantity: 8, AisleNo: "C23", Upc: "700000000000"},
			{Title: "Slot Machine", Price: 999.99, Quantity: 7, AisleNo: "C23", Upc: "800000000000"},
		}},
		{Title: "Groceries", Products: []Product{
			{Title: "Organic Apples (0.5 lb)", Price: 0.99, Quantity: 100, AisleNo: "A01", Upc: "900000000001"},
			{Title: "Whole Milk (1 gal)", Price: 3.49, Quantity: 50, AisleNo: "A01", Upc: "900000000002"},
			{Title: "Brown Eggs (Dozen)", Price: 2.99, Quantity: 60, AisleNo: "A01", Upc: "900000000003"},
			{Title: "Peanut Butter (16 oz)", Price: 2.59, Quantity: 40, AisleNo: "A01", Upc: "900000000004"},
			{Title: "Pasta (Spaghetti, 16 oz)", Price: 1.29, Quantity: 70, AisleNo: "A01", Upc: "900000000005"},
			{Title: "Canned Black Beans", Price: 0.89, Quantity: 80, AisleNo: "A01", Upc: "900000000006"},
			{Title: "Olive Oil (1L)", Price: 6.99, Quantity: 30, AisleNo: "A01", Upc: "900000000007"},
			{Title: "All-Purpose Flour (5 lb)", Price: 2.49, Quantity: 45, AisleNo: "A01", Upc: "900000000008"},
		}},
		{Title: "Toys", Products: []Product{
			{Title: "Lego Classic Bricks", Price: 29.99, Quantity: 25, AisleNo: "D12", Upc: "910000000001"},
			{Title: "Remote Control Car", Price: 49.99, Quantity: 20, AisleNo: "D12", Upc: "910000000002"},
			{Title: "Slinky", Price: 19.99, Quantity: 30, AisleNo: "D12", Upc: "910000000003"},
			{Title: "Board Game - Catan", Price: 44.99, Quantity: 15, AisleNo: "D12", Upc: "910000000004"},
			{Title: "Puzzle (1000 pieces)", Price: 14.99, Quantity: 25, AisleNo: "D12", Upc: "910000000005"},
			{Title: "Nerf Blaster", Price: 24.99, Quantity: 20, AisleNo: "D12", Upc: "910000000006"},
			{Title: "Barbie Dreamhouse Mini", Price: 39.99, Quantity: 18, AisleNo: "D12", Upc: "910000000007"},
			{Title: "Action Figure - Superhero", Price: 12.99, Quantity: 35, AisleNo: "D12", Upc: "910000000008"},
		}},
		{Title: "Office Supplies", Products: []Product{
			{Title: "Pencils (3-pack)", Price: 4.99, Quantity: 60, AisleNo: "E05", Upc: "920000000001"},
			{Title: "Notebook (200 pages)", Price: 3.49, Quantity: 50, AisleNo: "E05", Upc: "920000000002"},
			{Title: "Stapler", Price: 6.99, Quantity: 30, AisleNo: "E05", Upc: "920000000003"},
			{Title: "Printer Paper (500 sheets)", Price: 7.99, Quantity: 40, AisleNo: "E05", Upc: "920000000004"},
			{Title: "Desk Organizer", Price: 12.99, Quantity: 25, AisleNo: "E05", Upc: "920000000005"},
			{Title: "Highlighter Set (5 colors)", Price: 5.49, Quantity: 35, AisleNo: "E05", Upc: "920000000006"},
			{Title: "Sticky Notes (12-pack)", Price: 6.29, Quantity: 45, AisleNo: "E05", Upc: "920000000007"},
			{Title: "Whiteboard Markers (8-pack)", Price: 9.99, Quantity: 20, AisleNo: "E05", Upc: "920000000008"},
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

	routerGroup := router.Group("/api/browse")
	{
		routerGroup.GET("", getProductCategory)
	}

	router.Run()
}
