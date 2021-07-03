package ginrestaurant

import (
	"fooddelivery/module/restaurant/business"
	"fooddelivery/module/restaurant/model"
	"fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func CreateRestaurant(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var newRestaurant restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&newRestaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewCreateRestaurant(store)

		if err := business.CreateRestaurant(c.Request.Context(), &newRestaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": newRestaurant.ID})
	}
}
