package restaurantmodel

import (
	"fooddelivery/common"
)

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	OwnerId int    `json:"owner_id" gorm:"column:owner_id;"`
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}
