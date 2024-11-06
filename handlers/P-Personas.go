package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PutPersona(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		id := Request.Param("id")
		var Persona models.Persona
		if err := db.First(&Persona, id).Error; err != nil {
			Request.JSON(http.StatusNotFound, gin.H{"error": "Persona no encontrado " + err.Error()})
			return
		}

		var input models.Persona
		if err := Request.ShouldBindJSON(&input); err != nil {
			Request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Model(&Persona).Updates(input).Error; err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la Persona en la BD " + err.Error()})
			return
		}

		Request.JSON(http.StatusOK, Persona)
	}
}
