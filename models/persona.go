package models

import "gorm.io/gorm"

type Persona struct {
	gorm.Model

	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Objetos  []Objeto
}
