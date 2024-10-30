package database

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDBFromContext(ctx *gin.Context) *gorm.DB {
	database, exists := ctx.Get("db")
	if !exists {
		log.Fatalln("db context error")
		return nil
	}

	db, ok := database.(*gorm.DB)
	if !ok {
		log.Fatalln("db context error")
		return nil
	}

	return db
}
