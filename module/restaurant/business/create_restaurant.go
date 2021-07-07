package restaurantbusiness

import (
	"context"
	"errors"
	"fooddelivery/module/restaurant/model"
)

type CreateStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBusiness struct {
	store CreateStore
}

func NewCreateRestaurant(store CreateStore) *createRestaurantBusiness {
	return &createRestaurantBusiness{store: store}
}

func (biz createRestaurantBusiness) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("name cannot be null")
	}
	err := biz.store.Create(ctx, data)
	return err
}
