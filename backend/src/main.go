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
		log.Fatalln(err)
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

	adminGroup := router.Group("/admin", middleware.AdminKeyRequired(adminKey))
	adminGroup.GET("/clients", handlers.GetClients)
	adminGroup.POST("/clients/new", handlers.CreateClient)
	adminGroup.DELETE("/clients/:id", handlers.DeleteClient)
	adminGroup.GET("/cars", handlers.GetCars)
	adminGroup.POST("/cars/new", handlers.CreateCar)
	adminGroup.DELETE("/cars/:id", handlers.DeleteCar)
	adminGroup.GET("/rents", handlers.GetRents)
	adminGroup.POST("/rents/new", handlers.CreateRent)
	adminGroup.DELETE("/rents/:id", handlers.DeleteClient)

	router.POST("/register", handlers.CreateClient)
	router.GET("/cars", handlers.GetCars)
	router.POST("/order", handlers.CreateRent)
	router.GET("/orders", handlers.GetRentsByEmail)

	if err := router.Run(":8080"); err != nil {
		log.Fatalln(err)
		return
	}
}
