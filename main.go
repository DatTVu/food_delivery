package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

type SQLModel struct {
	ID        int       `gorm:"column:id;"`
	CreatedAt time.Time `gorm:"column:created_at;"`
	UpdatedAt time.Time `gorm:"column:updated_at;"`
	Status    int       `gorm:"column:status;default:1;"`
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
	OwnerId *int    `gorm:"column:owner_id;"`
	Name    *string `gorm:"column:name;"`
	Address *string `gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return "restaurants"
}

type RestaurantUpdate struct {
	SQLModel
	OwnerId *int    `gorm:"column:owner_id;"`
	Name    *string `gorm:"column:name;"`
	Address *string `gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return "restaurants"
}

func main() {
	log.Println("Hello world!")

	dsn := os.Getenv("CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("[err]MySql:", err)
	}

	db = db.Debug()

	var myNote Note
	if err := db.
		Where("id = ?", 7).
		First(&myNote).Error; err != nil {
		log.Println(err)
	}
	log.Println(myNote.Name, myNote.CategoryId)

	var listNote []Note
	if err := db.Find(&listNote).Error; err != nil {
		log.Println(err)
	}

	log.Println(listNote)
	myNote.Name = "Huong Mai"
	emptyString := ""
	zeroCategoryId := 0
	if err := db.
		Where("id=?", 9).
		Updates(NoteUpdate{
			Name:       &emptyString,
			CategoryId: &zeroCategoryId,
		}).Error; err != nil {
		log.Println(err)
	}

	log.Println(myNote)

	//Hard delete
	if err := db.
		Table(myNote.TableName()).
		Where("id=?", 9).
		Delete(nil); err != nil {
		log.Println(err)
	}

}
