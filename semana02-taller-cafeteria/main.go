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

//CRUD de cliente

func AgregarCliente(clientes []Cliente, nuevo Cliente) []Cliente {
	return append(clientes, nuevo)
}

func BuscarClientePorID(clientes []Cliente, id int) int {
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
	idx := BuscarClientePorID(clientes, id)
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

func BuscarProductoPorID(productos []Producto, id int) int {
	for i, t := range productos {
		if t.ID == id {
			return i
		}
	}
	return -1
}

func ListarProducto(productos []Producto) {
	fmt.Println("\n=== PRODUCTOS ===")
	if len(productos) == 0 {
		fmt.Println("(no hay prodcutos)")
		return
	}
	for _, p := range productos {
		fmt.Printf("[%d] %s | $%.2f | Stock: %d | Categoría: %s\n",
			p.ID, p.Nombre, p.Precio, p.Stock, p.Categoria)
	}
}

func EliminarProducto(productos []Producto, id int) []Producto {
	idx := BuscarProductoPorID(productos, id)
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

func leerEntero(lector *bufio.Reader, prompt string) int {
	fmt.Print(prompt)
	texto := leerLinea(lector)
	n, err := strconv.Atoi(texto)
	if err != nil {
		return -1
	}
	return n
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
		fmt.Println("2. Listar productos")
		fmt.Println("3. Agregar cleintes")
		fmt.Println("4. Agregar prodcutos")
		fmt.Print("Seleccione una opción: ")

		opcionStr := leerLinea(lector)

		switch opcionStr {
		case "1":
			ListarCliente(clientes)
		case "2":
			ListarProducto(productos)
		case "3": //Agregar cliente
			fmt.Println("\n--- Agregar cliente ---")
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			fmt.Print("Carrera: ")
			carrera := leerLinea(lector)
			fmt.Print("Saldo: ")
			saldoStr := leerLinea(lector)
			saldo, _ := strconv.ParseFloat(saldoStr, 64)
			nuevo := Cliente{id, nombre, carrera, saldo}
			clientes = AgregarCliente(clientes, nuevo)
			fmt.Println("✓ Cliente agregado correctamente")
		case "4": // Agregar prodcuto
			fmt.Println("\n--- Agregar producto ---")
			id := leerEntero(lector, "ID: ")
			fmt.Print("Nombre: ")
			nombre := leerLinea(lector)
			fmt.Print("Precio: ")
			precioStr := leerLinea(lector)
			precio, _ := strconv.ParseFloat(precioStr, 64)
			stock := leerEntero(lector, "Stock: ")
			fmt.Print("Categoría: ")
			categoria := leerLinea(lector)
			nuevo := Producto{id, nombre, precio, stock, categoria}
			productos = AgregarProducto(productos, nuevo)
			fmt.Println("✓ Producto agregado correctamente")
		}
	}
}
