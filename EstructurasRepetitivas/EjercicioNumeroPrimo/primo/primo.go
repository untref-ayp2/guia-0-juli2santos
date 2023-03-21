package primo

import "math"

// calculo si es primo con la forma de la raiz cuadrada

func EsNumeroPrimo(numero int) bool {

	// si es menor o igual a 1 no es
	if numero <= 1 {
		return false
	}

	// si la division por los menores de la raiz da entero tampoco es primo
	for i := 2; i <= int(math.Sqrt(float64(numero))); i++ {
		if numero%i == 0 {
			return false
		}
	}
	return true

}
