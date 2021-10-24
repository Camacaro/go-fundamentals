package main

import "fmt"

func main() {
	fmt.Println("============== Ejercicio 1 -  ==================")
}
package main

import "fmt"

func closeDb() {
	fmt.Println("close db")
}

func main() {

	fmt.Println("============== Ejecicio 1 - El uso de los keywords defer ==================")
	// Antes de morir el codigo se ejecuta el defer
	// Lo que hace es ejecutar la ultima funcion antes de que termine el programa
	// Case de uso que cierre la coneccion a la base de datos
	// Case de uso que un cierre el archivo al terminar el programa
	// Las buenas practicas dicen que solo se debe usar un defer por programa
	// La secuencia con mas de un defer es de la ultima a la primera

	// Antes de probrar el ejemplo de defer, comenta el otro ejercicio
	// defer fmt.Println("!!!")  // Se ejecuta de tercero
	// defer fmt.Println("Hola") // Se ejecuta segundo
	// fmt.Println("Mundo")      // Se ejecuta primero

	fmt.Println("============== Ejecicio 2 - El uso de los keywords continue y break ==================")
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		if i == 8 {
			fmt.Println("Break")
			break
		}
	}

fmt.Println("============== Reto - Close DB ==================")
defer closeDb()
fmt.Println("main")
fmt.Println("Se abre conexion db")
fmt.Println("Se guargo item en db")
fmt.Println("Se retorna valor")

}
