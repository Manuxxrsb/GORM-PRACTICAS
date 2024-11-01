package handlers

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Elimina_Persona(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		var Persona models.Persona
		if err := db.First(&Persona, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Persona not found"})
			return
		}

		if err := db.Delete(&Persona).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Persona"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Persona deleted successfully"})
	}
}

func Elimina_Objeto(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
			return
		}

		var Objeto models.Objeto
		if err := db.First(&Objeto, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Objeto not found"})
			return
		}

		if err := db.Delete(&Objeto).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Objeto"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Objeto deleted successfully"})
	}
}
