package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ndy-corp/1.src/midterm-1/src-code-new/initializers"
	"github.com/ndy-corp/1.src/midterm-1/src-code-new/models"
	"net/http"
	"strconv"
)

func CreateOrder(c *gin.Context) {
	var body struct {
		ProductID uint `json:"product_id"`
		Quantity  int  `json:"quantity"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	product := models.Product{}
	_ = initializers.DB.Find(&product, "id = ?", body.ProductID)

	order := models.Order{
		ProductID:  body.ProductID,
		Quantity:   body.Quantity,
		TotalPrice: product.Price * body.Quantity,
		Status:     models.IN_PROGRESS_STATUS,
	}

	result := initializers.DB.Create(&order)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create order",
		})

		return
	}

	c.JSON(http.StatusCreated, gin.H{"order": order})
}

func ChangeStatus(c *gin.Context) {
	params := c.Params

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return
	}

	order := models.Order{}

	_ = initializers.DB.Find(&order, "id = ?", id)

	if c.Request.Method == http.MethodDelete {
		order.Status = models.CANCEL_STATUS
	} else if c.Request.Method == http.MethodPatch {
		order.Status = models.PURCHASED_STATUS
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}
