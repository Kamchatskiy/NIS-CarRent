package database

import (
	"log"
	"os"
	"strings"

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
	dbPassStr := strings.TrimSpace(string(dbPass))

	dsn := "host=db user=backend password=" + dbPassStr + " dbname=car_rent port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	db.AutoMigrate(&models.Client{})
	db.AutoMigrate(&models.Rent{})
	db.AutoMigrate(&models.Car{})

	return db, nil
}
