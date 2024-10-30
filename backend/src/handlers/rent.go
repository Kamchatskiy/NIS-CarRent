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
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, rents)
}

func CreateRent(ctx *gin.Context) {
	db := database.GetDBFromContext(ctx)

	var rent models.Rent
	if err := ctx.ShouldBindJSON(&rent); err != nil {
		ctx.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	if !rent.EndDate.After(rent.StartDate) {
		ctx.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	var client models.Client
	err := db.Where("email = ?", rent.ClientEmail).First(&client).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			ctx.String(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
			return
		}
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	car, err := getCarByID(db, rent.CarID)
	if err != nil {
		ctx.String(http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	var tempRent models.Rent
	err = db.Where("car_id = ? AND ((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?))",
		rent.CarID, rent.EndDate, rent.StartDate, rent.StartDate, rent.EndDate).First(&tempRent).Error
	if err == nil {
		ctx.String(http.StatusConflict, http.StatusText(http.StatusConflict))
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	days := uint(rent.EndDate.Sub(rent.StartDate).Hours() / 24)
	rent.Price = car.DailyPrice*days + car.InsurancePrice

	if err := db.Create(&rent).Error; err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.String(http.StatusCreated, http.StatusText(http.StatusCreated))
}

func DeleteRent(ctx *gin.Context) {
	db := database.GetDBFromContext(ctx)

	rentID := ctx.Param("id")
	if rentID == "" {
		ctx.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
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
		return
	}

	ctx.String(http.StatusOK, http.StatusText(http.StatusOK))
}

func GetRentsByEmail(ctx *gin.Context) {
	db := database.GetDBFromContext(ctx)

	email := ctx.Query("email")
	if email == "" {
		ctx.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	client, err := getClientByEmail(db, email)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.JSON(http.StatusOK, client)
}
