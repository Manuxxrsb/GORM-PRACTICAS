package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePersona(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		var Persona models.Persona
		if err := Request.ShouldBindJSON(&Persona); err != nil {
			Request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&Persona).Error; err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}
		Request.JSON(http.StatusOK, Persona)
	}

}
