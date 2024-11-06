package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetallObjeto(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		var Objeto []models.Objeto
		err := db.Find(&Objeto).Error //todos los registros que coinciden con el modelo

		if err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"error": "No existen Peersonas " + err.Error()})
			return
		}
		Request.JSON(http.StatusOK, Objeto)
	}
}

func GetObjeto(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		id := Request.Param("id")
		var Objeto models.Objeto
		if err := db.First(&Objeto, id).Error; err != nil { //Primero encuentra el primer registro ordenado por clave principal
			Request.JSON(http.StatusNotFound, gin.H{"error": "Objeto no encontrado " + err.Error()})
			return
		}
		Request.JSON(http.StatusOK, Objeto)
	}
}
