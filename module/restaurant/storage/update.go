package restaurantstorage

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/module/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	if err := s.db.Where("id=?", id).
		Updates(data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}
		return common.ErrDB(err)
	}
	return nil
}
