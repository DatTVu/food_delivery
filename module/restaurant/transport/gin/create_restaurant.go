package ginrestaurant

import (
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	"fooddelivery/module/restaurant/business"
	"fooddelivery/module/restaurant/model"
	"fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appContext appctx.AppContext) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var newRestaurant restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&newRestaurant); err != nil {

			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}

		store := restaurantstorage.NewSQLStore(appContext.GetMainDBConnection())
		business := restaurantbusiness.NewCreateRestaurant(store)

		if err := business.CreateRestaurant(c.Request.Context(), &newRestaurant); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
			return
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(newRestaurant.ID))
	}
}
