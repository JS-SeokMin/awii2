package models

import "gorm.io/gorm"

// TAREA (CP1): Complete los campos de Compra según lo que muestran las pantallas.
//
// Pistas de trabajo:
//   - Un Compra referencia a una Funcion y a un Cliente (claves foráneas).
//   - Recuerde el campo de estado (use las constantes de estados.go) y el total.
//   - Los tests de acceptance/ compilan contra los nombres EXACTOS de los campos.
type Compra struct {
	gorm.Model
	FuncionID uint    `gorm:"not null" json:"funcion_id"`
	ClienteID uint    `gorm:"not null" json:"cliente_id"`
	Cantidad  uint    `gorm:"not null" json:"cantidad"`
	Estado    string  `gorm:"size:20;not null" json:"estado"`
	Total     float64 `gorm:"not null" json:"total"`
}
