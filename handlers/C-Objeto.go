package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateObjeto(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		var Objeto models.Objeto
		if err := Request.ShouldBindJSON(&Objeto); err != nil {
			Request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&Objeto).Error; err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		Request.JSON(http.StatusOK, Objeto)
	}

}
