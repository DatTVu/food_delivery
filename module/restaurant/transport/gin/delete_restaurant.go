package ginrestaurant

import (
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	restaurantbusiness "fooddelivery/module/restaurant/business"
	restaurantstorage "fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func DeleteRestaurant(appContext appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
		business := restaurantbusiness.NewDeleteRestaurantBusiness(store)

		if err_ := business.DeleteRestaurant(c.Request.Context(), id); err_ != nil {
			c.JSON(http.StatusInternalServerError, common.ErrInternal(err_))
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
		return
	}
}
