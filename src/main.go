package main

import "fmt"

func main() {
	fmt.Println("======== Ejercicio 1 =========")
	// Declaración de constantes - Nunca va a cambiar de valor
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("Hola Mundo")
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	fmt.Println("======== Ejercicio 2 =========")

	// Declaracion de variables enteras
	base := 12 // la crea y asigna el tipo de dato con los :
	var altura int = 14
	var area int

	fmt.Println("base:", base)
	fmt.Println("altura:", altura)
	fmt.Println("area:", area)

	fmt.Println("======== Ejercicio 3 =========")
	// Zero values - Valores por defecto
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println("Enteros:", a)
	fmt.Println("float64:", b)
	fmt.Println("Strings:", c) // Es un string vacio
	fmt.Println("Booleanas:", d)

	fmt.Println("======== Ejercicio 4 =========")
	// Area Cuadradp
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado es:", areaCuadrado)

	fmt.Println("======== Ejercicio 5 - Operadores arimeticos =========")
	x := 10
	y := 50
	result := x + y
	fmt.Println("Suma:", result)

	result = y - x
	fmt.Println("Resta:", result)
}
