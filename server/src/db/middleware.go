package database

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MiddleWareDB(ctx *gin.Context) *gorm.DB {
	database, exists := ctx.Get("db")
	if !exists {
		log.Fatalln("db context error")
		return nil
	}
	db, _ := database.(*gorm.DB)
	return db
}
