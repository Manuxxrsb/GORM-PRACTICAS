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

	db.AutoMigrate(models.Objeto{})
	db.AutoMigrate(models.Persona{})

	//Personas Routes
	router.POST("/Persona", handlers.CreatePersona(db))
	router.GET("/Personas", handlers.GetallPersona(db))
	router.GET("/Persona/:id", handlers.GetPersona(db))
	router.PUT("/Persona/:id", handlers.PutPersona(db))
	router.DELETE("/Persona/:id", handlers.Delete_Persona(db))

	//Objetos Routes
	router.POST("/Objeto", handlers.CreateObjeto(db))
	router.GET("/Objetos", handlers.GetallObjeto(db))
	router.GET("/Objeto/:id", handlers.GetObjeto(db))
	router.PUT("/Objeto/:id", handlers.PutObjeto(db))
	router.DELETE("/Objeto/:id", handlers.Delete_Objeto(db))

}
