package main

import (
	"fmt"
	"go-fundamentals/greetings"
	pk "go-fundamentals/mypackage"
)

func main() {
	// En los modificadores de visibilidad, podemos usar la primer letra en mayúscula para indicar que es de acceso público
	// y en minúscula para indicar que es de acceso privado.
	fmt.Println("============== Ejercicio 1 - Modificadores de acceso en funciones y Structs  ==================")

	fmt.Println(" ============== Package Greetings")
	greetings.Hello()

	fmt.Println(" ============== Package Mypackage")
	var myCar pk.CarPublic
	myCar.Brand = "Ford"
	myCar.Year = 2020
	fmt.Println("myCar", myCar)

	// No se puede acceder ya que es privado
	// var myOtherCar pk.carPrivate
	// fmt.Println("myOtherCar", myOtherCar)

	fmt.Println(" ============== Package Mypackage, PrintMessage")
	pk.PrintMessage("Jesus")
}
