package main

import (
	"fooddelivery/component/appctx"
	ginrestaurant "fooddelivery/module/restaurant/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type SQLModel struct {
	ID        int       `json:"id" gorm:"column:id;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;"`
	Status    int       `json:"status" gorm:"column:status;default:1;"`
}

type Restaurant struct {
	SQLModel
	OwnerId int    `gorm:"column:owner_id;"`
	Name    string `gorm:"column:name;"`
	Address string `gorm:"column:addr;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	log.Println("Hello world!")

	dsn := os.Getenv("CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("[err]MySql:", err)
	}

	db = db.Debug()

	appContext := appctx.NewAppContext(db)

	r := gin.Default()
	v1 := r.Group("/v1")

	{
		restaurants := v1.Group("/restaurants")
		{
			//CREATE
			restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))
			//GET
			restaurants.GET("", ginrestaurant.ListRestaurant(appContext))
			//GET ID
			restaurants.GET("/:id", ginrestaurant.GetRestaurant(appContext))
			//UPDATE
			restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(appContext))
			//DELETE
			restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
		}
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
