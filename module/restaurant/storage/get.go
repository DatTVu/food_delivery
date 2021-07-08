package restaurantstorage

import (
	"context"
	"fooddelivery/common"
	"fooddelivery/module/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) GetDataWithCondition(ctx context.Context, condition map[string]interface{}) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
