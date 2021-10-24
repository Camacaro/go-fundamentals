package pc

import "fmt"

// Fijarte en las mayuscula de las palabras ya que indican que son publicas

type Pc struct {
	Ram   int
	Dick  int
	Brand string
}

// Agregar funciones a los structs
func (myPc Pc) Ping() {
	fmt.Println("myPc", myPc.Brand, "Pong")
}

// Aqui lo que hacemos es acceder a los valores del struct y cambiar el valor de la ram
// Lo hacemos con el puntero para acceder a su espacio de memoria donde esta el valor
func (myPc *Pc) DuplicateRAM() {
	myPc.Ram = myPc.Ram * 2
}
