package handlers

import (
	"log"
	"net/http"

	database "github.com/Kamchatskiy/NIS-CarRent/db"
	"github.com/Kamchatskiy/NIS-CarRent/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateClient(ctx *gin.Context) {
	db := database.GetDBFromContext(ctx)

	var client models.Client
	if err := ctx.ShouldBindJSON(&client); err != nil {
		log.Println(err)
		ctx.String(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		ctx.Abort()
		return
	}

	var tempClient models.Client
	if err := db.Where("email = ?", client.Email).First(&tempClient).Error; err == nil {
		ctx.String(http.StatusConflict, "client already exists")
		ctx.Abort()
		return
	} else if err != gorm.ErrRecordNotFound {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		ctx.Abort()
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
	db := database.GetDBFromContext(ctx)

	var clients []models.Client
	if err := db.Find(&clients).Error; err != nil {
		log.Println(err)
		ctx.String(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		ctx.Abort()
		return
	}
	ctx.JSON(http.StatusOK, clients)
}
