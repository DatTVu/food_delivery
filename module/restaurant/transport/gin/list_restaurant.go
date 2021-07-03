package ginrestaurant

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/module/restaurant/business"
	"fooddelivery/module/restaurant/model"
	restaurantstorage "fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type fakeListStore struct{}

func (fakeListStore) ListDataWithCondition(ctx context.Context, filter *restaurantmodel.Filter, paging *common.Paging) ([]restaurantmodel.Restaurant, error) {
	return []restaurantmodel.Restaurant{
		{
			SQLModel: common.SQLModel{ID: 1},
			Name:     "AA",
			Address:  "BB",
		},
	}, nil
}

func ListRestaurant(db *gorm.DB) func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		paging.Process()

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		business := restaurantbusiness.NewlistRestaurantBusiness(store)

		result, err := business.ListRestaurant(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result, "paging": paging, "filter": filter})
	}
}
