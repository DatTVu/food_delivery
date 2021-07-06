package main

import (
	ginrestaurant "fooddelivery/module/restaurant/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"strconv"
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

	r := gin.Default()
	v1 := r.Group("/v1")

	{
		restaurants := v1.Group("/restaurants")
		{
			//CREATE
			restaurants.POST("", ginrestaurant.CreateRestaurant(db))
			//GET
			restaurants.GET("", ginrestaurant.ListRestaurant(db))
			//GET ID
			restaurants.GET("/:id", ginrestaurant.GetRestaurant(db))
			//UPDATE
			restaurants.PUT("/:id", ginrestaurant.UpdateRestaurant(db))
			//DELETE
			restaurants.DELETE("/:id", func(c *gin.Context) {
				id_, err_ := strconv.Atoi(c.Param("id"))
				if err_ != nil {
					c.JSON(http.StatusBadRequest, gin.H{"data": err})
					return
				}

				if err := db.
					Table(Restaurant{}.TableName()).
					Where("id = ?", id_).
					Delete(nil).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": 1})
				return
			})
		}
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
