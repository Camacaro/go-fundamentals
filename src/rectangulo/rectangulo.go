package rectangulo

type Reactangulo struct {
	Base   float64
	Altura float64
}

// Methods
func (r Reactangulo) Area() float64 {
	return r.Base * r.Altura
}
