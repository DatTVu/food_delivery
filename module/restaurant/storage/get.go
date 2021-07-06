package restaurantstorage

import (
	"context"
	"fooddelivery/module/restaurant/model"
)

func (s *sqlStore) GetDataWithCondition(ctx context.Context, condition map[string]interface{}) (*restaurantmodel.Restaurant, error) {
	var data restaurantmodel.Restaurant
	if err := s.db.Where(condition).First(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
