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
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
		business := restaurantbusiness.NewgetRestaurantBusiness(store)

		data, err_ := business.GetRestaurant(c.Request.Context(), id)

		if err_ != nil {
			panic(err_)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
		return
	}
}
