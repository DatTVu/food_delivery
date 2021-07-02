package main

import (
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

type Note struct {
	SQLModel
	Name       string `gorm:"column:title;"`
	CategoryId int    `gorm:"column:category_id;"`
}

func (Note) TableName() string {
	return "notes"
}

type NoteUpdate struct {
	Name       *string `gorm:"column:title;"`
	CategoryId *int    `gorm:"column:category_id;"`
	Status     *int    `gorm:"column:status;"`
}

func (NoteUpdate) TableName() string {
	return Note{}.TableName()
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

type RestaurantCreate struct {
	SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
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
			//CRUD Block
			//CREATE
			restaurants.POST("", func(c *gin.Context) {
				var newRestaurant RestaurantCreate
				if err := c.ShouldBind(&newRestaurant); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err})
					return
				}
				if err := db.Create(&newRestaurant).Error; err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}
				c.JSON(http.StatusOK, gin.H{"data": newRestaurant.ID})
			})

			//GET
			restaurants.GET("", func(c *gin.Context) {
				var restaurantsList []Restaurant

				var paging struct {
					Page  int   `json:"page" form:"page"`
					Limit int   `json:"limit" form:"limit"`
					Total int64 `json:"total" form:"total "`
				}

				if err := c.ShouldBind(&paging); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if paging.Limit <= 0 {
					paging.Limit = 10
				}

				if paging.Page <= 0 {
					paging.Page = 1
				}

				if err := db.Table(Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				if err := db.
					Limit(paging.Limit).
					Offset((paging.Page - 1) * paging.Limit).
					Order("id desc").
					Find(&restaurantsList).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": restaurantsList, "paging": paging})

			})

			//GET ID
			restaurants.GET("/:id", func(c *gin.Context) {
				id_, err_ := strconv.Atoi(c.Param("id"))
				if err_ != nil {
					c.JSON(http.StatusBadRequest, gin.H{"data": err})
					return
				}

				var restaurant Restaurant

				if err := db.
					Where("id = ?", id_).
					First(&restaurant).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": restaurant})
				return
			})

			//UPDATE
			restaurants.PUT("/:id", func(c *gin.Context) {
				id_, err_ := strconv.Atoi(c.Param("id"))
				if err_ != nil {
					c.JSON(http.StatusBadRequest, gin.H{"data": err})
					return
				}

				var updatedData RestaurantUpdate

				if err := c.ShouldBind(&updatedData); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err})
					return
				}

				if err := db.
					Table(updatedData.TableName()).
					Where("id = ?", id_).
					Updates(updatedData).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
					return
				}

				c.JSON(http.StatusOK, gin.H{"data": 1})
				return
			})

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
