package database

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDBFromContext(ctx *gin.Context) *gorm.DB {
	database, exists := ctx.Get("db")
	if !exists {
		log.Fatalln("db context error: database not found in context")
		return nil
	}

	db, ok := database.(*gorm.DB)
	if !ok {
		log.Fatalln("db context error: database not found in context")
		return nil
	}

	return db
}
