package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

// ===== CRUD CLIENTE =====

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
		fmt.Printf("⚠ El cliente con el ID %d no existe.\n", id)
		return clientes
	}
	return append(clientes[:idx], clientes[idx+1:]...)
}

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

func RegistrarPedido(
	clientes []Cliente,
	productos []Producto,
	pedidos []Pedido,
	clienteID int,
	productoID int,
	cantidad int,
) ([]Pedido, error) {

	idxC := BuscarClienteID(clientes, clienteID)
	if idxC == -1 {
		return pedidos, fmt.Errorf("cliente no existe")
	}

	idxP := BuscarProductoID(productos, productoID)
	if idxP == -1 {
		return pedidos, fmt.Errorf("producto no existe")
	}

	if productos[idxP].Stock < cantidad {
		return pedidos, fmt.Errorf("stock insuficiente")
	}

	total := float64(cantidad) * productos[idxP].Precio

	if clientes[idxC].Saldo < total {
		return pedidos, fmt.Errorf("saldo insuficiente")
	}

	productos[idxP].Stock -= cantidad
	clientes[idxC].Saldo -= total

	nuevo := Pedido{
		ID:         len(pedidos) + 1,
		ClienteID:  clienteID,
		ProductoID: productoID,
		Cantidad:   cantidad,
		Total:      total,
	}

	pedidos = append(pedidos, nuevo)
	return pedidos, nil
}

func leerLinea(lector *bufio.Reader) string {
	linea, _ := lector.ReadString('\n')
	return strings.TrimSpace(linea)
}

func leerEntero(lector *bufio.Reader, msg string) int {
	fmt.Print(msg)
	texto := leerLinea(lector)
	n, _ := strconv.Atoi(texto)
	return n
}

func leerFloat(lector *bufio.Reader, msg string) float64 {
	fmt.Print(msg)
	texto := leerLinea(lector)
	f, _ := strconv.ParseFloat(texto, 64)
	return f
}

func main() {

	clientes := []Cliente{
		{1, "Juan Mora", "Medicina", 100},
		{2, "Steven Delgado", "Software", 50},
		{3, "José Anchundia", "Arquitectura", 21},
	}

	productos := []Producto{
		{1, "Energizante Volt", 1.50, 20, "bebida"},
		{2, "Deditos de queso", 1.25, 15, "snack"},
		{3, "Arroz con pollo", 2.50, 10, "almuerzo"},
	}

	pedidos := []Pedido{}

	lector := bufio.NewReader(os.Stdin)

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
		fmt.Println("9. Registrar pedido")

		fmt.Print("Seleccione una opción: ")

		opcionStr := leerLinea(lector)

		switch opcionStr {

		case "1":
			ListarCliente(clientes)

		case "2":
			id := leerEntero(lector, "ID cliente: ")
			idx := BuscarClienteID(clientes, id)
			if idx == -1 {
				fmt.Println("Cliente no encontrado")
			} else {
				fmt.Println(clientes[idx])
			}

		case "3":
			fmt.Println("\n--- Agregar Cliente ---")
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			fmt.Print("Carrera: ")
			carrera := leerLinea(lector)
			saldo := leerFloat(lector, "Saldo: ")

			clientes = AgregarCliente(clientes, Cliente{id, nombre, carrera, saldo})

		case "4":
			id := leerEntero(lector, "ID a eliminar: ")
			clientes = EliminarCliente(clientes, id)

		case "5":
			ListarProductos(productos)

		case "6":
			id := leerEntero(lector, "ID producto: ")
			idx := BuscarProductoID(productos, id)
			if idx == -1 {
				fmt.Println("Producto no encontrado")
			} else {
				fmt.Println(productos[idx])
			}

		case "7":
			fmt.Println("\n--- Agregar Producto ---")
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			precio := leerFloat(lector, "Precio: ")
			stock := leerEntero(lector, "Stock: ")
			fmt.Print("Categoría (bebida/snack/almuerzo): ")
			cat := leerLinea(lector)

			if cat != "bebida" && cat != "snack" && cat != "almuerzo" {
				fmt.Println("Categoría inválida")
				continue
			}

			productos = AgregarProducto(productos, Producto{id, nombre, precio, stock, cat})

		case "8":
			id := leerEntero(lector, "ID a eliminar: ")
			productos = EliminarProducto(productos, id)

		case "9":
			idC := leerEntero(lector, "ID Cliente: ")
			idP := leerEntero(lector, "ID Producto: ")
			cant := leerEntero(lector, "Cantidad: ")

			var err error
			pedidos, err = RegistrarPedido(clientes, productos, pedidos, idC, idP, cant)

			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Pedido registrado")
			}

		default:
			fmt.Println("Opción inválida")
		}
	}
}
