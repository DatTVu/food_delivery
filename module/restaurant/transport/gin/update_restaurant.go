package ginrestaurant

import (
	"fooddelivery/module/restaurant/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func UpdateRestaurant(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err})
			return
		}

		var updatedData restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&updatedData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		if err := db.
			Table(updatedData.TableName()).
			Where("id = ?", id).
			Updates(updatedData).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": 1})
		return
	}
}
