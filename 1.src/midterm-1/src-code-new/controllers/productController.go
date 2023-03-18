package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ndy-corp/1.src/midterm-1/src-code-new/initializers"
	"github.com/ndy-corp/1.src/midterm-1/src-code-new/models"
	"net/http"
)

func ListProducts(c *gin.Context) {
	qs := c.Request.URL.Query()
	s := qs.Get("title")
	var products []models.Product
	_ = initializers.DB.Find(&products)
	if s != "" {
		_ = initializers.DB.Find(&products, "title = ?", s)
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func CreateProduct(c *gin.Context) {
	var body struct {
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	product := models.Product{Title: body.Title, Description: body.Description}
	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create product",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}
