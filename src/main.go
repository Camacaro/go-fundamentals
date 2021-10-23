package main

import "fmt"

func main() {

	fmt.Println("============== Ejecicio 1 - For Condicional ==================")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("============== Ejecicio 2 - For While ==================")
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	fmt.Println("============== Ejecicio 3 - For forever ==================")
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
		if counterForever == 10 {
			break
		}
	}

	fmt.Println("============== Ejecicio 4 - For Range  ==================")
	listaNumerosPares := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	for i, par := range listaNumerosPares {
		fmt.Printf("posicion %d número par: %d \n", i, par)
	}

	fmt.Println("============== Reto  ==================")
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}
