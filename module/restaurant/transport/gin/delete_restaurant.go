package ginrestaurant

import (
	restaurantbusiness "fooddelivery/module/restaurant/business"
	restaurantstorage "fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func DeleteRestaurant(db *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"data": err})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewDeleteRestaurantBusiness(store)

		if err_ := business.DeleteRestaurant(c.Request.Context(), id); err_ != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"data": err_.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": 1})
		return
	}
}
