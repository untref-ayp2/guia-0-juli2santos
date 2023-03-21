package main

import (
	"fmt"
	"producto/producto"
)

func main() {

	fmt.Println("ingrese el primer valor")
	var valor int
	fmt.Scanf("%d", &valor)

	fmt.Println("ingrese el segundo valor")
	var valor2 int
	fmt.Scanf("%d", &valor2)

	total := producto.Multiplicar(valor, valor2)

	fmt.Println(total)

}
