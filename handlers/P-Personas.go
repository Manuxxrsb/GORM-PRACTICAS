package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PutPersona(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {

		//Obtener el ID de usuario del cuerpo de la solicitud
		var Persona_Nueva models.Persona
		if err := Request.BindJSON(&Persona_Nueva); err != nil {
			Request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Encuentra al usuario en la base de datos
		var Persona models.Persona
		db.First(&Persona, Persona.ID)
		if Persona.ID == 0 {
			Request.JSON(http.StatusNotFound, gin.H{"error": "Persona no encontrado"})
			return
		}
		// Actualiza al usuario en la base de datos
		err := db.Model(&Persona).Updates(Persona_Nueva).Error
		if err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar en la BD" + err.Error()})
			return
		}

		Request.JSON(http.StatusOK, gin.H{"message": "Persona actualizado con éxito"})
	}
}
