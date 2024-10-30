package database

import (
	"log"
	"os"

	"github.com/Kamchatskiy/NIS-CarRent/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dbPass, err := os.ReadFile("/run/secrets/db-pass")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dsn := "host=db user=backend password=" + string(dbPass) + " dbname=car_rent port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	db.AutoMigrate(&models.Rent{})
	db.AutoMigrate(&models.Client{})
	db.AutoMigrate(&models.Car{})

	return db, nil
}
