package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ndy-corp/1.src/midterm-1/src-code-new/initializers"
	"github.com/ndy-corp/1.src/midterm-1/src-code-new/models"
	"net/http"
	"strconv"
)

func CreateComment(c *gin.Context) {
	var body struct {
		Body   string
		Rating float32
	}
	params := c.Params

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	user, _ := c.Get("user")

	usr, ok := user.(models.User)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to assertion to type",
		})
		return
	}
	//usr := &models.User{}
	var product models.Product
	initializers.DB.First(&product, "id = ?", id)
	var sum float32
	var cnt float32
	comment := models.Comment{Body: body.Body, Rating: body.Rating, User: usr.ID, Product: int(id)}

	result := initializers.DB.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create comment",
		})

		return
	}

	initializers.DB.Table("comments").Where("product = ?", id).Select("sum(rating) as sm").Scan(&sum)
	initializers.DB.Table("comments").Where("product = ?", id).Select("count(*) as cnt").Scan(&cnt)
	avg := sum / cnt
	initializers.DB.Model(&product).Update("rating", avg)

	c.JSON(http.StatusOK, gin.H{"comment": comment})
}
