package handlers

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete_Persona(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id " + err.Error()})
			return
		}

		var Persona models.Persona
		if err := db.First(&Persona, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No se pudo encontrar la persona para eliminar " + err.Error()})
			return
		}

		if err := db.Delete(&Persona).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la Persona " + err.Error()})
			return
		}
		//delete no elimina la tabla pero para gorm llena deletedat que lo ignorarara para los find o first

		c.JSON(http.StatusOK, gin.H{"message": "Persona eliminado con éxito "})
	}
}
