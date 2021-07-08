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
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		var updatedData restaurantmodel.RestaurantUpdate

		if err := c.ShouldBind(&updatedData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
		business := restaurantbusiness.NewUpdateRestaurantBusiness(store)

		if err := business.UpdateRestaurant(c, id, &updatedData); err != nil {
			c.JSON(http.StatusInternalServerError, common.ErrInternal(err))
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(1))
		return
	}
}
