package main

import (
	"fmt"
	modulePC "go-fundamentals/src/pc"
)

func main() {
	fmt.Println("============== Ejercicio 1 - Structs y Punteros  ==================")
	a := 50
	// Apuntar b hacia el espacio de memoria de a
	// &a direccion de memoria de a
	b := &a

	fmt.Println("Valor de a:", a)      // 50
	fmt.Println("Dirección de a:", &a) // 0xc0000bc000
	fmt.Println("Valor de b:", b)      // 0xc0000bc000
	// con *b accedemos al valor de b
	fmt.Println("Valor de b:", *b) // 50

	*b = 100
	// Cambia el valor de a, ya que ambos apuntan al mismo espacio de memoria
	fmt.Println("Valor de a:", a) // 100

	// Se usa muchos los punteros hacer funciones mas customizadas
	// y para pasar llevvar funciones hacias otros paquetes mucho mas eficiente

	fmt.Println("============== Ejercicio 2 - Funciones en Structs  ==================")
	myPc := modulePC.Pc{Ram: 16, Dick: 200, Brand: "Dell"}
	fmt.Println("Valor de myPc:", myPc)
	myPc.Ping()

	// Con esto podemos cambiar el valor de la ram, pero si el programa es muy grande estarias redundando
	// myPc.ram = 20
	myPc.DuplicateRAM()
	fmt.Println("Valor de mi nueva ram:", myPc)

	myPc.DuplicateRAM()
	fmt.Println("Valor de mi nueva ram 2:", myPc)
}

// ============== Antes del Refactor =================
// ============== Modulo pc  				 =================
/*
	package main

	import "fmt"

	type pc struct {
		ram   int
		dick  int
		brand string
	}

	// Agregar funciones a los structs
	func (myPc pc) ping() {
		fmt.Println("myPc", myPc.brand, "Pong")
	}

	// Aqui lo que hacemos es acceder a los valores del struct y cambiar el valor de la ram
	// Lo hacemos con el puntero para acceder a su espacio de memoria donde esta el valor
	func (myPc *pc) duplicateRAM() {
		myPc.ram = myPc.ram * 2
	}

	func main() {
		fmt.Println("============== Ejercicio 1 - Structs y Punteros  ==================")
		a := 50
		// Apuntar b hacia el espacio de memoria de a
		// &a direccion de memoria de a
		b := &a

		fmt.Println("Valor de a:", a)      // 50
		fmt.Println("Dirección de a:", &a) // 0xc0000bc000
		fmt.Println("Valor de b:", b)      // 0xc0000bc000
		// con *b accedemos al valor de b
		fmt.Println("Valor de b:", *b) // 50

		*b = 100
		// Cambia el valor de a, ya que ambos apuntan al mismo espacio de memoria
		fmt.Println("Valor de a:", 100) // 100

		// Se usa muchos los punteros hacer funciones mas customizadas
		// y para pasar llevvar funciones hacias otros paquetes mucho mas eficiente

		fmt.Println("============== Ejercicio 2 - Funciones en Structs  ==================")
		myPc := pc{ram: 16, dick: 200, brand: "Dell"}
		fmt.Println("Valor de myPc:", myPc)
		myPc.ping()

		// Con esto podemos cambiar el valor de la ram, pero si el programa es muy grande estarias redundando
		// myPc.ram = 20
		myPc.duplicateRAM()
		fmt.Println("Valor de mi nueva ram:", myPc)

		myPc.duplicateRAM()
		fmt.Println("Valor de mi nueva ram 2:", myPc)
	}
*/
