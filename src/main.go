package main

import "fmt"

func main() {

	fmt.Println("============== Ejercicio 1 - Recorrido de Slices con Range ==================")
	slice := []string{"Hola", "que", "tal", "?"}

	for indice, value := range slice {
		fmt.Println(indice, value)
	}
}
