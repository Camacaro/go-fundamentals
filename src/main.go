package main

import "fmt"

type figure2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type reactangulo struct {
	base   float64
	altura float64
}

// El nombre de las funciones no se pueden repetir en un mismos paquete,
// pero que esta pasando con las func area que se definieron dos veces en
// esta clase?
// Lo que esta sucediendo es que no son simplemente funciones,
// en Go se les conoce como Methods. Un método es una función con un
// argumento de receptor especial. que debe de ser de tipo struct o
// interface y se definen de la siguiente manera:
/*
	func (receiver type)  nameFunc() {
		// body del metodo
	}
*/

// Methods
func (c cuadrado) area() float64 {
	return c.base * c.base
}

// Methods
func (r reactangulo) area() float64 {
	return r.base * r.altura
}

func calcularArea(figura figure2D) {
	fmt.Println("El area es: ", figura.area())
}

func main() {
	fmt.Println("========== Ejercicio 1 - Interfaces ==========")
	myCuadraro := cuadrado{base: 5}
	myReactangulo := reactangulo{base: 2, altura: 4}

	// De Esta menera se puede llamar los metodos de area de cada uno de los struct

	// Si los structs que tenemos en el código tienen métodos que hacen algo
	// en común (Cálculos, obtener data, etc), es posible ejecutar éstos
	// métodos usando una interfaz, de esta forma evitamos hacer código
	// por cada struct.

	calcularArea(myCuadraro)
	calcularArea(myReactangulo)

	// De esta  foma llamas directamente el .area() del struct
	fmt.Println(figure2D.area(myCuadraro))
	fmt.Println(figure2D.area(myReactangulo))

	fmt.Println("========== Ejercicio 2 - listas de interfaces ==========")
	myInterface := []interface{}{"Hola", 12, 4.90, true}
	fmt.Println("Destructurado")
	fmt.Println(myInterface...)
	fmt.Println("Array o slide tipo interfaz", myInterface)

	fmt.Println("========== Ejercicio 3 - range at interfaces ==========")
	for _, valor := range myInterface {
		fmt.Println(valor)
	}
}
