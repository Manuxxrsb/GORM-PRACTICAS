package handlers

import (
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Delete_Objeto(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id " + err.Error()})
			return
		}

		var Objeto models.Objeto
		if err := db.First(&Objeto, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "No se pudo encontrar la persona para eliminar " + err.Error()})
			return
		}

		if err := db.Delete(&Objeto).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la Objeto " + err.Error()})
			return
		}
		//delete no elimina la tabla pero para gorm llena deletedat que lo ignorarara para los find o first

		c.JSON(http.StatusOK, gin.H{"message": "Objeto eliminado con éxito "})
	}
}
