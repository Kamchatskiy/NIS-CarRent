package handlers

import (
	"log"
	"net/http"

	database "github.com/Kamchatskiy/NIS-CarRent/db"
	"github.com/Kamchatskiy/NIS-CarRent/models"
	"github.com/gin-gonic/gin"
)

func GetCars(ctx *gin.Context) {
	db := database.GetDBFromContext(ctx)

	var cars []models.Car
	if err := db.Find(&cars).Error; err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		ctx.Abort()
		return
	}

	ctx.JSON(http.StatusOK, cars)
}
