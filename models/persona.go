package models

type Persona struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Nombre   string
	Apellido string
	Objetos  []Objeto `gorm:"foreignKey:PersonaID;references:ID"`
}

type Objeto struct {
	ID        uint `gorm:"primaryKey;autoIncrement"`
	Nombre    string
	PersonaID uint `gorm:"references:ID"`
}
