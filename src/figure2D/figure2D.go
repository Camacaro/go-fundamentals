package figure2D

import "fmt"

type figure2D interface {
	Area() float64
}

func CalcularArea(figura figure2D) {
	fmt.Println("El area es: ", figura.Area())
}
