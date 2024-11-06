package models

import "gorm.io/gorm"

type Objeto struct {
	gorm.Model

	Nombre    string `json:"nombre"`
	PersonaID uint   `json:"persona_id"`
}
