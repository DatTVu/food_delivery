package restaurantstorage

import (
	"context"
	"fooddelivery/module/restaurant/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := s.db.Where("id=?", id).
		Updates(data).Error; err != nil {
		return err
	}
	return nil
}
