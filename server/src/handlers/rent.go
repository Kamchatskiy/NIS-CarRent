package handlers

import (
	"log"
	"net/http"

	database "github.com/Kamchatskiy/NIS-CarRent/db"
	"github.com/Kamchatskiy/NIS-CarRent/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func DeleteRent(ctx *gin.Context) {
	db := database.GetDBFromContext(ctx)

	rentID := ctx.Param("id")
	if rentID == "" {
		ctx.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		ctx.Abort()
		return
	}

	var rent models.Rent
	if err := db.First(&rent, rentID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		} else {
			log.Println(err)
			ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		ctx.Abort()
		return
	}

	if err := db.Delete(&rent).Error; err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		ctx.Abort()
		return
	}

	ctx.String(http.StatusOK, http.StatusText(http.StatusOK))
}
