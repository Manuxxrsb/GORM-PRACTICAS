package handlers

import (
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func PutObjeto(db *gorm.DB) gin.HandlerFunc {
	return func(Request *gin.Context) {

		//Obtener el ID de usuario del cuerpo de la solicitud
		var Objeto_nuevo models.Objeto
		if err := Request.BindJSON(&Objeto_nuevo); err != nil {
			Request.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Encuentra al usuario en la base de datos
		var Objeto models.Objeto
		db.First(&Objeto, Objeto.ID)
		if Objeto.ID == 0 {
			Request.JSON(http.StatusNotFound, gin.H{"error": "Objeto no encontrado"})
			return
		}

		// Actualiza al usuario en la base de datos
		err := db.Model(&Objeto).Updates(Objeto_nuevo).Error
		if err != nil {
			Request.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar en la BD" + err.Error()})
			return
		}

		Request.JSON(http.StatusOK, gin.H{"message": "Objeto actualizado con éxito"})
	}
}
