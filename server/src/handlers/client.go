package handlers

import (
	"log"
	"net/http"

	database "github.com/Kamchatskiy/NIS-CarRent/db"
	"github.com/Kamchatskiy/NIS-CarRent/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context) {
	db := database.MiddleWareDB(ctx)

	var client models.Client
	if err := ctx.ShouldBindJSON(&client); err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	var tempClient models.Client
	if err := db.Where("email = ?", client.Email).First(&tempClient).Error; err == nil {
		ctx.String(http.StatusConflict, "client already exists")
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	if err := db.Create(&client).Error; err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}

	ctx.String(http.StatusOK, http.StatusText(http.StatusOK))
}

func GetClients(ctx *gin.Context) {
	db := database.MiddleWareDB(ctx)

	if err := ctx.ShouldBindHeader("Admin-Key"); err != nil {
		log.Println(err)
		ctx.String(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	key := ctx.GetHeader("Admin-Key")
	adminKey, exists := ctx.Get("admin-key")
	if !exists {
		log.Fatalln("admin-key context error")
		return
	}
	if adminKey != key {
		ctx.String(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		return
	}

	var clients []models.Client
	if err := db.Find(&clients); err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	ctx.JSON(http.StatusOK, clients)
}
