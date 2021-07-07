package restaurantbusiness

import (
	"context"
	"errors"
	"fooddelivery/module/restaurant/model"
)

type GetRestaurantStore interface {
	GetDataWithCondition(ctx context.Context, condition map[string]interface{}) (*restaurantmodel.Restaurant, error)
}

type getRestaurantBusiness struct {
	store GetRestaurantStore
}

func NewgetRestaurantBusiness(store GetRestaurantStore) *getRestaurantBusiness {
	return &getRestaurantBusiness{store: store}
}

func (business *getRestaurantBusiness) GetRestaurant(ctx context.Context,
	id int,
) (*restaurantmodel.Restaurant, error) {
	data, err := business.store.GetDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	if data.Status == 0 {
		return nil, errors.New("restaurant has been deleted!")
	}

	return data, nil
}
