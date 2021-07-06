package restaurantmodel

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
	Status  *int    `json:"-" gorm:"column:status;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
