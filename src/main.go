package main

import (
	"fmt"
	fibonacci "go-fundamentals/src/fibonacci"
	myclass "go-fundamentals/src/myclass"
	observable "go-fundamentals/src/observable"
)

func main() {
	fmt.Println("======= Ejercicio de fibonacci =======")
	fibonacci.Main()

	fmt.Println("======= Ejercicio Mi clase =======")
	myclass.Main()

	fmt.Println("======= Ejercicio de Observador - Observable: =======")
	observable.Main()
}
