package main

import "fmt"

func main() {
	fmt.Println("============ Ejercicio 1 - Println ============")
	// Declaracion de variable
	helloMessage := "Hello"
	worldMessage := "Word"

	// Println: Agrega una linea
	fmt.Println(helloMessage, worldMessage)

	fmt.Println("============ Ejercicio 2 - Printf ============")
	// Printf
	nombre := "Platzi"
	cursos := 500

	// %s para indicar que ahi va una variable de string
	// %d para indicar que ahi va una variable de entero
	fmt.Printf("%s tiene más de %d cursos \n", nombre, cursos)

	// si no se que tipo de varible va puedo colocar v - Lo recomendable es que si sabes el tipo de datos lo uses
	fmt.Printf("%v tiene más de %v cursos \n", nombre, cursos)

	// Usamos %T para determinar el tipo de dato de una variable a utilizar.
	fmt.Printf("%T\n", nombre)

	fmt.Println("============ Ejercicio 3 - Sprintf ============")
	// Con este Sprintf podemos generar un string y lo guardamos no lo va a imprimir en consola
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println("message:", message)
}
