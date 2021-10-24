package main

import "fmt"

func isPar(n int) bool {
	return n%2 == 0
}

func login(user, password string) bool {
	return user == "admin" && password == "admin"
}

func main() {

	fmt.Println("============== Ejecicio 1 - if ==================")
	valor1 := 1
	valor2 := 2

	if valor1 == valor2 {
		fmt.Println("Son iguales")
	} else {
		fmt.Println("No son iguales")
	}

	fmt.Println("============== Ejecicio 2 - Operadores logico AND  ==================")
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("Son iguales")
	} else {
		fmt.Println("No son iguales")
	}

	fmt.Println("============== Ejecicio 3 - Operadores logico OR  ==================")
	if valor1 == 0 || valor2 == 2 {
		fmt.Println("Son iguales")
	} else {
		fmt.Println("No son iguales")
	}

	fmt.Println("============== Reto 1 - Numero Par  ==================")
	numero := 10
	if isPar(numero) {
		fmt.Println("El numero es par")
	} else {
		fmt.Println("El numero es impar")
	}

	fmt.Println("============== Reto 1 - Login  ==================")
	user := "admin"
	password := "admin"
	if login(user, password) {
		fmt.Println("Login correcto")
	} else {
		fmt.Println("Login incorrecto")
	}
}
