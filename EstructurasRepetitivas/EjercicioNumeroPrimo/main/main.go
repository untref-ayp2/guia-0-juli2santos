package main

import (
	"fmt"
	"primo/primo"
)

func main() {

	var valor int
	fmt.Scanf("%d", &valor)

	primo := primo.EsNumeroPrimo(valor)

	fmt.Println(primo)

}
