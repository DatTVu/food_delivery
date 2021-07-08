package restaurantstorage

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/module/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {
	if err := s.db.
		Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id=?", id).
		Delete(nil).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return common.RecordNotFound
		}
		return common.ErrDB(err)
	}
	return nil
}
