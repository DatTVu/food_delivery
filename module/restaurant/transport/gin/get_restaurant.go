package ginrestaurant

import (
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	"fooddelivery/module/restaurant/business"
	"fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
		business := restaurantbusiness.NewgetRestaurantBusiness(store)

		data, err_ := business.GetRestaurant(c.Request.Context(), id)

		if err_ != nil {
			c.JSON(http.StatusBadRequest, err_)
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
		return
	}
}
