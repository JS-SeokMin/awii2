package storage

import (
	"gorm.io/gorm"

	"github.com/joancema/examen-cine/internal/models"
)

// TAREA (CP2): Implemente CompraGORM contra la interfaz CompraRepository.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Guíese por FuncionGORM: es el mismo patrón con una entidad distinta.
//   - Recuerde: aquí NO va lógica de negocio. Solo persistencia.
type CompraGORM struct {
	db *gorm.DB
}

func NuevaCompraGORM(db *gorm.DB) *CompraGORM {
	return &CompraGORM{db: db}
}

func (r *CompraGORM) Crear(a *models.Compra) error {
	return r.db.Create(a).Error
}

func (r *CompraGORM) ObtenerPorID(id uint) (models.Compra, bool) {
	var a models.Compra
	if err := r.db.First(&a, id).Error; err != nil {
		return models.Compra{}, false
	}
	return a, true
}

func (r *CompraGORM) Listar() ([]models.Compra, error) {
	var lista []models.Compra
	err := r.db.Find(&lista).Error
	return lista, err
}

func (r *CompraGORM) Actualizar(a *models.Compra) error {
	return r.db.Save(a).Error
}
