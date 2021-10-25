package main

import "fmt"

type pc struct {
	ram   int
	brand string
	disk  int
}

// La estructura de datos " Struct " tiene un método llamado
// " String " , que podemos sobrescribir para personalizar
// la salida a consola de los datos del struct.
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 8, brand: "Dell", disk: 500}

	// Se imprimira de otra forma
	fmt.Println(myPc)
}
