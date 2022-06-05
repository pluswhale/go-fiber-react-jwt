package database

import (
	"github.com/pluswhale/go-auth-react/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	//connect to DB
	dsn := "host=192.168.0.109 user=postgres password=123 dbname=yt_authorization port=5432 sslmode=disable TimeZone=Europe/Moscow"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&model.User{})
}
