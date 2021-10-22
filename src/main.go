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
}
