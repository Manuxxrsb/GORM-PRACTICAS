package routes

import (
	"backend/handlers"
	"backend/models"
	"backend/utils"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func ConfiguraRutas(router *gin.Engine) {

	db, err := utils.OpenGormDB() //abro la conexion a la base de datos
	if err != nil {
		log.Fatalf("Error al conectarse a la BD: %v", err)
	}

	fmt.Print(db)

	db.AutoMigrate(&models.Persona{}, &models.Objeto{})

	//Create
	router.POST("/Persona", handlers.CreatePersona(db))
	router.POST("/Objeto", handlers.CreateObjeto(db))

	//Read
	router.GET("/Personas", handlers.GetallPersona(db))
	router.GET("/Persona/:id", handlers.GetPersona(db))
	router.GET("/Objetos", handlers.GetallObjeto(db))
	router.GET("/Objeto/:id", handlers.GetObjeto(db))
	//Delete
	router.DELETE("/Persona/:id", handlers.Elimina_Persona(db))
	router.DELETE("/Objeto/:id", handlers.Elimina_Objeto(db))

}
