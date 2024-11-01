package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPersona(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		id := informacion.Param("Id")
		var Persona models.Persona
		if err := db.First(&Persona, id).Error; err != nil {
			informacion.JSON(http.StatusNotFound, gin.H{"error": "Persona no encontrado"})
			return
		}
		informacion.JSON(http.StatusOK, Persona)
	}
}

func GetallPersona(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		var Persona []models.Persona
		err := db.Find(&Persona).Error
		if err != nil {
			informacion.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar todos los Persona"})
			return
		}
		informacion.JSON(http.StatusOK, Persona)
	}
}

func GetObjeto(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		id := informacion.Param("Id")
		var Objeto models.Objeto
		if err := db.First(&Objeto, id).Error; err != nil {
			informacion.JSON(http.StatusNotFound, gin.H{"error": "Objeto no encontrado"})
			return
		}
		informacion.JSON(http.StatusOK, Objeto)
	}
}

func GetallObjeto(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		var Objeto []models.Objeto
		err := db.Find(&Objeto).Error
		if err != nil {
			informacion.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar todos los Objeto"})
			return
		}
		informacion.JSON(http.StatusOK, Objeto)
	}
}
