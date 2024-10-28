package handlers

import (
	"log"
	"net/http"

	database "github.com/Kamchatskiy/NIS-CarRent/db"
	"github.com/Kamchatskiy/NIS-CarRent/models"
	"github.com/gin-gonic/gin"
)

func GetRents(ctx *gin.Context) {
	db := database.GetDBFromContext(ctx)

	var rents []models.Rent
	if err := db.Find(&rents).Error; err != nil {
		log.Fatalln(err)
		return
	}

	ctx.JSON(http.StatusOK, rents)
}
