package trapecio

import "fmt"

type trapecio struct {
	baseSmall, baseTall, altura float64
}

func New(baseSmall, baseTall, altura float64) *trapecio {
	return &trapecio{baseSmall, baseTall, altura}
}

func (r trapecio) String() string {
	return fmt.Sprintf("Trapecio en su base pqueña = %b, base grande = %b y altura %b", r.baseSmall, r.baseTall, r.altura)
}

func (t *trapecio) Area() float64 {
	return ((t.baseSmall + t.baseTall) / 2) * t.altura
}
