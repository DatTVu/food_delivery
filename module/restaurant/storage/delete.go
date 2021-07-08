package restaurantstorage

import (
	"context"
	"fooddelivery/module/restaurant/model"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id=?", id).
		Delete(nil).Error; err != nil {
		return err
	}
	return nil
}
