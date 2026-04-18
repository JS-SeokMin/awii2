package main

import "fmt"

func main() {

	var age int
	fmt.Println("Enter your age:")
	fmt.Scanln(&age)
	//go conditionals
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// go switch
	var diaDeLaSemana int = 1
	switch diaDeLaSemana {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// different types of for

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for i := range 10 {
		fmt.Println(i)
	}
	for i, value := range []int{1, 2, 3, 4, 5} {
		fmt.Println(value)
		fmt.Println(i)
	}
	for _, value := range []int{1, 2, 3, 4, 5} {
		fmt.Println(value)
	}
	numeros := []int{1, 2, 33, 4, 50}
	for _, numero := range numeros {
		fmt.Println(numero)
	}

	for {
		fmt.Println("Infinite loop")
		break
	}
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

}
