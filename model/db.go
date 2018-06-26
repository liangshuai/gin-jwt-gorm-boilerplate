package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func InitDB(connection string) {
	var err error
	db, err = gorm.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
}

func GetDB() *gorm.DB {
	return db
}

func AutoMigrate() {
	db.AutoMigrate(&User{}, &Post{}, &Comment{}, &Tag{}, &Correlation{})
}