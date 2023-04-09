package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ndy-corp/1.src/midterm-1/src-code-new/initializers"
	"github.com/ndy-corp/1.src/midterm-1/src-code-new/models"
	"net/http"
)

func ListProducts(c *gin.Context) {
	qs := c.Request.URL.Query()
	inputTitle := qs.Get("title")
	inputSort := qs.Get("sort")
	var products []models.Product
	if inputTitle != "" && inputSort != "" {
		if inputSort[0] == '-' {
			_ = initializers.DB.Order(fmt.Sprint(inputSort[1:])).Find(&products, "title = ?", inputTitle)
		} else {
			_ = initializers.DB.Order(fmt.Sprint(inputSort, " desc")).Find(&products, "title = ?", inputTitle)
		}
	} else if inputTitle != "" {
		_ = initializers.DB.Find(&products, "title = ?", inputTitle)
	} else if inputSort != "" {
		if inputSort[0] == '-' {
			_ = initializers.DB.Order(fmt.Sprint(inputSort[1:])).Find(&products)
		} else {
			_ = initializers.DB.Order(fmt.Sprint(inputSort, " desc")).Find(&products)
		}
	} else {
		_ = initializers.DB.Find(&products)
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func CreateProduct(c *gin.Context) {
	var body struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
		Price       int    `json:"price,omitempty"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	product := models.Product{Title: body.Title, Description: body.Description, Price: body.Price}
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create product",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
