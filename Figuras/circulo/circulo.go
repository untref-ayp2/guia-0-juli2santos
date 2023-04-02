package circulo

import (
	punto "Figuras/Punto"
	"fmt"
	"math"
)

type Circulo struct {
	P punto.Punto
	R float32
}

func (c Circulo) Area() float32 {
	return c.R * c.R * math.Pi
}

func (c Circulo) Perimetro() float32 {
	return c.R * 2 * math.Pi
}

func (c Circulo) ToString() string {
	cadena := fmt.Sprint("Circulo:", c)

	return cadena
}
