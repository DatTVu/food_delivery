package restaurantbusiness

import (
	"context"
	"errors"
	restaurantmodel "fooddelivery/module/restaurant/model"
)

type DeleteStore interface {
	GetDataWithCondition(ctx context.Context, condition map[string]interface{}) (*restaurantmodel.Restaurant, error)
	Delete(ctx context.Context, id int) error
}

type deleteRestaurantBusiness struct {
	store DeleteStore
}

func NewDeleteRestaurantBusiness(store DeleteStore) *deleteRestaurantBusiness {
	return &deleteRestaurantBusiness{store: store}
}

func (business *deleteRestaurantBusiness) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := business.store.GetDataWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("Restaurant has been deleted!")
	}
	if err := business.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
