package cuadrado

type Cuadrado struct {
	Base float64
}

func (c Cuadrado) Area() float64 {
	return c.Base * c.Base
}
