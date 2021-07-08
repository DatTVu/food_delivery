package restaurantbusiness

import (
	"context"
	"errors"
	"fooddelivery/module/restaurant/model"
)

type UpdateStore interface {
	GetDataWithCondition(ctx context.Context, condition map[string]interface{}) (*restaurantmodel.Restaurant, error)
	Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBusiness struct {
	store UpdateStore
}

func NewUpdateRestaurantBusiness(store UpdateStore) *updateRestaurantBusiness {
	return &updateRestaurantBusiness{store: store}
}

func (business *updateRestaurantBusiness) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	oldData, err := business.store.GetDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("Restaurant has been deleted!")
	}
	if err := business.store.Update(ctx, id, data); err != nil {
		return err
	}
	return nil
}
