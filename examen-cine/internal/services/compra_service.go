package services

import (
	"github.com/joancema/examen-cine/internal/models"
	"github.com/joancema/examen-cine/internal/storage"
)

// TAREA (CP2): Implemente CompraService con las 5 reglas de negocio.
//
// Las reglas están A LA VISTA en las pantallas (carpeta pantallas/) y los
// tests de acceptance/reglas_test.go las verifican una por una. Devuelva los
// errores de dominio de errores.go: los tests los comprueban con errors.Is.
//
// Reglas:
//   - NO cambie el nombre del tipo, del constructor ni las firmas de los métodos.
//   - Observe que el service recibe TRES repositories: necesita consultar
//     Funcion y Cliente para validar, y actualizar Funcion al cancelar.
type CompraService struct {
	compras   storage.CompraRepository
	funciones storage.FuncionRepository
	clientes  storage.ClienteRepository
}

func NuevaCompraService(
	compras storage.CompraRepository,
	funciones storage.FuncionRepository,
	clientes storage.ClienteRepository,
) *CompraService {
	return &CompraService{
		compras:   compras,
		funciones: funciones,
		clientes:  clientes,
	}
}

// Crear registra un nuevo compra aplicando R1, R2 y R3.
// TODO (R1): la funcion debe existir y estar activa; el cliente debe existir.
// TODO (R2): la cantidad no puede superar el stock disponible de la funcion.
// TODO (R3): calcule el total (observe en las pantallas cuándo aplica descuento).
// TODO: al crear, el stock de la funcion se descuenta (mire la pantalla 01
// antes y después de crear una compra; R5 es la operación inversa).
func (s *CompraService) Crear(a *models.Compra) error {
	// R1: Función existe y está activa
	funcion, ok := s.funciones.ObtenerPorID(a.FuncionID)
	if !ok || !funcion.Activo {
		return ErrReferenciaInvalida
	}

	// R1: Cliente existe
	_, ok = s.clientes.ObtenerPorID(a.ClienteID)
	if !ok {
		return ErrReferenciaInvalida
	}

	// R2: Cantidad no supera el stock
	if a.Cantidad > funcion.Stock {
		return ErrStockInsuficiente
	}

	// R3: Calcular total con descuento si es >= 5 unidades
	total := float64(a.Cantidad) * funcion.PrecioUnitario
	if a.Cantidad >= 5 {
		total = total * 0.9 // 10% de descuento
	}
	a.Total = total
	a.Estado = models.EstadoPendiente

	// Crear la compra
	if err := s.compras.Crear(a); err != nil {
		return err
	}

	// R5: Descontar stock de la función
	funcion.Stock -= a.Cantidad
	if err := s.funciones.Actualizar(&funcion); err != nil {
		return err
	}

	return nil
}

func (s *CompraService) ObtenerPorID(id uint) (models.Compra, error) {
	c, ok := s.compras.ObtenerPorID(id)
	if !ok {
		return models.Compra{}, ErrNoEncontrado
	}
	return c, nil
}

func (s *CompraService) Listar() ([]models.Compra, error) {
	return s.compras.Listar()
}

// Cancelar cancela una compra aplicando R4 y R5.
// TODO (R4): solo se puede cancelar una compra en estado PENDIENTE.
// TODO (R5): al cancelar, la cantidad se repone al stock de la funcion.
func (s *CompraService) Cancelar(id uint) error {
	// Obtener la compra
	compra, ok := s.compras.ObtenerPorID(id)
	if !ok {
		return ErrNoEncontrado
	}

	// R4: Solo se puede cancelar si está PENDIENTE
	if compra.Estado != models.EstadoPendiente {
		return ErrEstadoInvalido
	}

	// R5: Reponer stock a la función
	funcion, ok := s.funciones.ObtenerPorID(compra.FuncionID)
	if !ok {
		return ErrNoEncontrado
	}

	funcion.Stock += compra.Cantidad
	if err := s.funciones.Actualizar(&funcion); err != nil {
		return err
	}

	// Cambiar estado a CANCELADA
	compra.Estado = models.EstadoCancelada
	return s.compras.Actualizar(&compra)
}
