package ginrestaurant

import (
	restaurantbusiness "fooddelivery/module/restaurant/business"
	"fooddelivery/module/restaurant/model"
	restaurantstorage "fooddelivery/module/restaurant/storage"
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

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewUpdateRestaurantBusiness(store)

		if err := business.UpdateRestaurant(c, id, &updatedData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": 1})
		return
	}
}
