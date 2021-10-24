// package main

// import "fmt"

// func main() {

// 	fmt.Println("============== Ejecicio 1 - Switch ==================")
// 	modulo := 4 % 2
// 	switch modulo {
// 	case 0:
// 		fmt.Println("El modulo es 0 - Es par")
// 	default:
// 		fmt.Println("El modulo es 1 - default - Es impar")
// 	}

// 	fmt.Println("============== Ejecicio 2 - Parceo de variables ==================")
// 	// Define una variable y luego la parceo
// 	switch modulo2 := 4 % 2; modulo2 {
// 	case 0:
// 		fmt.Println("El modulo es 0 - Es par")
// 	default:
// 		fmt.Println("El modulo es 1 - default - Es impar")
// 	}

// 	fmt.Println("============== Ejecicio 3 - Sin condición ==================")
// 	// Este caso es util cuando no se necesita una condición
// 	// Y se vaya a iterar sobre una misma variable
// 	// ó haya muchas condiciones
// 	value := 40
// 	switch {
// 	case value > 100:
// 		fmt.Println("El valor es mayor a 100")
// 	case value > 50:
// 		fmt.Println("El valor es mayor a 50")
// 	case value < 0:
// 		fmt.Println("El valor es menor a 0")
// 	default:
// 		fmt.Println("No condicion - default")
// 	}
// }
