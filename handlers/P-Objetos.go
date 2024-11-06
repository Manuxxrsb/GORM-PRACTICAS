package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PutObjeto(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		id := Request.Param("id")
		var Objeto models.Objeto
		if err := db.First(&Objeto, id).Error; err != nil {
			Request.JSON(http.StatusNotFound, gin.H{"error": "Objeto no encontrado " + err.Error()})
			return
		}

		var input models.Objeto
		if err := Request.ShouldBindJSON(&input); err != nil {
			Request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Model(&Objeto).Updates(input).Error; err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la Objeto en la BD " + err.Error()})
			return
		}

		Request.JSON(http.StatusOK, Objeto)
	}
}
