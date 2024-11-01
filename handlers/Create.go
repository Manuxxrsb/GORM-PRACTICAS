package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePersona(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		var Persona models.Persona
		if err := informacion.ShouldBindJSON(&Persona); err != nil {
			informacion.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for _, objeto := range Persona.Objetos {
			if err := db.Create(&objeto).Error; err != nil {
				informacion.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		if err := db.Create(&Persona).Error; err != nil {
			informacion.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		informacion.JSON(http.StatusOK, Persona)
	}
}

func CreateObjeto(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		var Objeto models.Objeto
		if err := informacion.ShouldBindJSON(&Objeto); err != nil { //verificamos info
			informacion.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := db.Create(&Objeto).Error; err != nil { // creamos la Objeto en la bd y controlamos el error
			informacion.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		informacion.JSON(http.StatusOK, Objeto)
	}
}
