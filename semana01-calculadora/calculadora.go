package main

import "fmt"

func main() {

	fmt.Println("==== CALCULADORA CIENTÍFICA v1.0 ====")

	// //Reton 1
	// const a = 40
	// const b = 30

	// const suma = a+b
	// const resta = a-b
	// const multi = a*b
	// const divi = a/b

	

	// fmt.Println("La suma total es:", suma, "\nResta total: ", resta, "\nMultiplicación total: ", multi, "\nDivisión: ", divi)

	//Retro 

	
	

	var historial string
	var contador int


	for { 

		var a, b float64
		var operacion string
		var resultado float64
		var linea string

		
		fmt.Print("\nIngresa el primer número: ")
		fmt.Scan(&a)

		fmt.Print("Ingresa el segundo número: ")
		fmt.Scan(&b)

		fmt.Print("Ingresa la operación (+, -, *, /, ^, !): ")
		fmt.Scan(&operacion)

		switch operacion {

		case "+":
			resultado = a + b
			linea = fmt.Sprintf("%.2f + %.2f = %.2f", a, b, resultado)

		case "-":
			resultado = a - b
			linea = fmt.Sprintf("%.2f - %.2f = %.2f", a, b, resultado)

		case "*":
			resultado = a * b
			linea = fmt.Sprintf("%.2f * %.2f = %.2f", a, b, resultado)

		case "/":
			if b == 0 {
				fmt.Println("Error: no se puede dividir entre cero")
				continue 
			}
			resultado = a / b
			linea = fmt.Sprintf("%.2f / %.2f = %.2f", a, b, resultado)

		case "^":
			if b < 0 {
				fmt.Println("Error: el exponente debe ser entero positivo")
				continue
			}

			exp := int(b)
			resultado = 1

			for i := 0; i < exp; i++ {
				resultado *= a
			}

			linea = fmt.Sprintf("%.2f ^ %d = %.2f", a, exp, resultado)

		case "!":
			if a < 0 {
				fmt.Println("Error: no existe factorial de números negativos")
				continue
			}

			n := int(a)
			if a != float64(n) {
				fmt.Println("Error: el factorial solo se define para enteros")
				continue
			}

			resultado = 1
			for i := 1; i <= n; i++ {
				resultado *= float64(i)
			}

			linea = fmt.Sprintf("%d ! = %.0f", n, resultado)

		default:
			fmt.Println("Error: operación no reconocida")
			continue
		}

		
		fmt.Println("Resultado:", linea)

		
		contador++
		historial = historial + fmt.Sprintf("%d) %s\n", contador, linea)

	    // ¿Continuar?
		var opcion string
		fmt.Print("¿Otra operación? (s/n): ")
		fmt.Scan(&opcion)

		if opcion == "n" || opcion == "N" {
			break 
		}
	}

	// Historial
	fmt.Println("\n==== HISTORIAL DE OPERACIONES ====")
	fmt.Print(historial)
	fmt.Printf("Total de operaciones realizadas: %d\n", contador)
}






