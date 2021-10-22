package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

// Normal( a int, b int, c string) -> Buena Practica (a, b int, c string)
func tripeArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func areaRectangulo(base, altura int) int {
	return base * altura
}

func areaTrapecio(baseSmall, baseTall, altura int) int {
	return altura * ((baseSmall + baseTall) / 2)
}

func areaCirculo(radio int) float64 {
	const pi float64 = 3.141516
	return pi * float64((radio * radio))
}

func main() {
	normalFunction("Hola Mundo")
	tripeArgument(1, 2, "Hola")

	value := returnValue(2)
	fmt.Println("value", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("doubleReturn - value1, value2 -", value1, value2)

	value11, _ := doubleReturn(2)
	fmt.Println("doubleReturn - value11", value11)

	_, value22 := doubleReturn(2)
	fmt.Println("doubleReturn - value22", value22)

	fmt.Println("=============== Reto ====================")
	rectangulo := areaRectangulo(10, 5)
	trapecio := areaTrapecio(3, 6, 4)
	circulo := areaCirculo(3)

	fmt.Printf("Area del rectangulo: %d, Area del trapecio: %d, Area del circulo: %b \n", rectangulo, trapecio, circulo)
}
