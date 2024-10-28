package main

import (
	"log"
	"os"

	database "github.com/Kamchatskiy/NIS-CarRent/db"
	"github.com/Kamchatskiy/NIS-CarRent/handlers"
	"github.com/Kamchatskiy/NIS-CarRent/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalln("error connecting to db")
		return
	}
	defer func() {
		tempDB, err := db.DB()
		if err != nil {
			log.Fatalln(err)
			return
		}
		if err := tempDB.Close(); err != nil {
			log.Fatalln(err)
			return
		}
	}()

	adminKeyBytes, err := os.ReadFile("/run/secrets/admin-key")
	if err != nil {
		log.Fatalln(err)
		return
	}
	adminKey := string(adminKeyBytes)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(func(ctx *gin.Context) {
		ctx.Set("db", db)
		ctx.Next()
	})

	router.GET("/clients", middleware.AdminKeyRequired(adminKey), handlers.GetClients)
	router.POST("/clients/new", handlers.CreateClient)

	router.GET("/cars", handlers.GetCars)
	router.POST("/cars/new", middleware.AdminKeyRequired(adminKey), handlers.CreateCar)

	router.GET("/rents", middleware.AdminKeyRequired(adminKey), handlers.GetRents)

	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err)
		return
	}
}
