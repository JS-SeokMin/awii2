package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Cliente struct {
	ID      int
	Nombre  string
	Carrera string
	Saldo   float64
}

type Producto struct {
	ID        int
	Nombre    string
	Precio    float64
	Stock     int
	Categoria string
}

type Pedido struct {
	ID         int
	ClienteID  int
	ProductoID int
	Cantidad   int
	Total      float64
	Fecha      string
}

//CRUD de cliente

func AgregarCliente(clientes []Cliente, nuevo Cliente) []Cliente {
	return append(clientes, nuevo)
}

func BuscarClienteID(clientes []Cliente, id int) int {
	for i, t := range clientes {
		if t.ID == id {
			return i
		}
	}
	return -1
}

func ListarCliente(clientes []Cliente) {
	fmt.Println("\n=== CLIENTES REGISTRADOS ===")

	if len(clientes) == 0 {
		fmt.Println("(no hay clientes)")
		return
	}

	for _, c := range clientes {
		fmt.Printf("[%d] %s | %s | Saldo: $%.2f\n",
			c.ID, c.Nombre, c.Carrera, c.Saldo)
	}
}

func EliminarCliente(clientes []Cliente, id int) []Cliente {
	idx := BuscarClienteID(clientes, id)
	if idx == -1 {
		fmt.Printf("⚠ Cliente con ID %d no existe.\n", id)
		return clientes
	}
	return append(clientes[:idx], clientes[idx+1:]...)
}

//CRUD de Producto

func AgregarProducto(productos []Producto, nuevo Producto) []Producto {
	return append(productos, nuevo)
}

func BuscarProductoID(productos []Producto, id int) int {
	for i, t := range productos {
		if t.ID == id {
			return i
		}
	}
	return -1
}

func ListarProductos(productos []Producto) {
	fmt.Println("\n=== PRODUCTOS ===")
	for _, p := range productos {
		fmt.Printf("[%d] %s | $%.2f | Stock: %d | Categoría: %s\n",
			p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria)
	}
}

func EliminarProducto(productos []Producto, id int) []Producto {
	idx := BuscarProductoID(productos, id)
	if idx == -1 {
		fmt.Printf("⚠ Producto con ID %d no existe.\n", id)
		return productos
	}
	return append(productos[:idx], productos[idx+1:]...)
}

func leerLinea(lector *bufio.Reader) string {
	linea, _ := lector.ReadString('\n')
	return strings.TrimSpace(linea)
}

func main() {
	clientes := []Cliente{
		{1, "Juan Mora", "Medicina", 100},
		{2, "Steven Delgado", "Software", 50},
		{3, "José Anchundia", "Arquitectura", 21},
	}
	_ = clientes

	productos := []Producto{
		{1, "Arroz marinero", 2.50, 10, "Comida"},
		{2, "Coca cola", 0.50, 34, "Bebidas"},
		{3, "Encebollado", 2, 25, "Comida"},
		{4, "Sopa de pollo", 2.15, 20, "Sopa"},
	}
	_ = productos

	pedidos := []Pedido{}
	_ = pedidos

	var lector = bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MENÚ ===")
		fmt.Println("1. Listar clientes")
		fmt.Println("2. Buscar clientes")
		fmt.Println("3. Agregar clientes")
		fmt.Println("4. Eliminar clientes")

		fmt.Println("5. Listar productos")
		fmt.Println("6. Buscar productos")
		fmt.Println("7. Agregar productos")
		fmt.Println("8. Eliminar productos")

		fmt.Print("Seleccione una opción: ")

		opcionStr := leerLinea(lector)

		switch opcionStr {
		case "1":
			ListarCliente(clientes)
		case "2":
			BuscarClienteID()
			fmt.Print()
		case "3":
			AgregarCliente()

			fmt.Println("\n--- Agregar Cliente ---")
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			fmt.Print("Tipo: ")
			tipo := leerLinea(lector)
			fmt.Print("Carrera: ")
			carrera := leerLinea(lector)
			cupos := leerEntero(lector, "Cupos disponibles: ")
			costo := leerFloat(lector, "Costo por visita: ")
			fmt.Print("Idiomas (separados por coma, ej: español,inglés): ")
			idiomasRaw := leerLinea(lector)
			idiomas := strings.Split(idiomasRaw, ",")
			for i := range idiomas {
				idiomas[i] = strings.TrimSpace(idiomas[i])
			}
			negocios = AgregarNegocio(negocios, Negocio{
				id, nombre, tipo, idiomas, ciudad, cupos, costo,
			})
			fmt.Println("✓ Negocio agregado.")

		case "4":
			EliminarCliente()
		case "5":
			ListarProductos()
		case "6":
			BuscarProducto()
		case "7":
			AgregarProducto()
		case "8":
			EliminarProducto()

		default:
			fmt.Println("Opción inválida")
		}
	}
}
