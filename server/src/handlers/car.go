package handlers

import (
	"log"
	"net/http"

	database "github.com/Kamchatskiy/NIS-CarRent/db"
	"github.com/Kamchatskiy/NIS-CarRent/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func CreateCar(ctx *gin.Context) {
	db := database.GetDBFromContext(ctx)

	var car models.Car
	if err := ctx.ShouldBindJSON(&car); err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		ctx.Abort()
		return
	}

	var tempCar models.Car
	if err := db.Where("model = ?", car.Model).First(&tempCar).Error; err == nil {
		ctx.String(http.StatusConflict, http.StatusText(http.StatusConflict))
		ctx.Abort()
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		ctx.Abort()
		return
	}

	if err := db.Create(&car).Error; err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.String(http.StatusCreated, http.StatusText(http.StatusCreated))
}

func DeleteCar(ctx *gin.Context) {
	db := database.GetDBFromContext(ctx)

	carID := ctx.Param("id")
	if carID == "" {
		ctx.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		ctx.Abort()
		return
	}

	var car models.Car
	if err := db.First(&car, carID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		} else {
			log.Println(err)
			ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		}
		ctx.Abort()
		return
	}

	if err := db.Delete(&car).Error; err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		ctx.Abort()
		return
	}

	ctx.String(http.StatusOK, http.StatusText(http.StatusOK))
}
