// Nombre del paquete (nombre de la carpeta) mypackage
package mypackage

import "fmt"

// En los modificadores de visibilidad, podemos usar la primer letra en mayúscula para indicar que es de acceso público
// y en minúscula para indicar que es de acceso privado.

// CarPublic es una variable de acceso público
type CarPublic struct {
	Brand string
	Year  int
}

type carPrivate struct {
	brand string
	year  int
}

// PrintMessage imprimirá un mensaje
func PrintMessage(txt string) {
	fmt.Println("Hello world", txt)
}

type User struct {
	Name string
	Age  int
}
