package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetallPersona(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		var Persona []models.Persona
		err := db.Find(&Persona).Error //todos los registros que coinciden con el modelo

		if err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"error": "No existen Peersonas " + err.Error()})
			return
		}
		Request.JSON(http.StatusOK, Persona)
	}
}

func GetPersona(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {
		id := Request.Param("id")
		var Persona models.Persona
		if err := db.First(&Persona, id).Error; err != nil { //Primero encuentra el primer registro ordenado por clave principal
			Request.JSON(http.StatusNotFound, gin.H{"error": "Persona no encontrado " + err.Error()})
			return
		}

		//agregamos la busqueda para los objetos de la Persona
		db.Model(&Persona).Association("Objetos").Find(&Persona.Objetos)
		Request.JSON(http.StatusOK, Persona)
	}
}
