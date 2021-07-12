package ginrestaurant

import (
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	restaurantbusiness "fooddelivery/module/restaurant/business"
	"fooddelivery/module/restaurant/model"
	restaurantstorage "fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UpdateRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var updatedData restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&updatedData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
		business := restaurantbusiness.NewUpdateRestaurantBusiness(store)

		if err := business.UpdateRestaurant(c, id, &updatedData); err != nil {
			panic(common.ErrInternal(err))
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
		return
	}
}
