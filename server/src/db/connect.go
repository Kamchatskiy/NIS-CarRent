package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	dbPass, err := os.ReadFile("/run/secrets/db-pass")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dsn := "host=db user=backend password=" + string(dbPass) + "dbname=car_rent port=9920 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}
