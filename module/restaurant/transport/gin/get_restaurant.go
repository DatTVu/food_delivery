package ginrestaurant

import (
	"fooddelivery/module/restaurant/business"
	"fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetRestaurant(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewgetRestaurantBusiness(store)

		data, err := business.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data})
		return
	}
}
