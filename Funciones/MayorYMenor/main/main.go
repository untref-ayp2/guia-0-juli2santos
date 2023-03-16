package main

import (
	"fmt"
	"lista/lista"
)

func main() {

	list := []int{5, 2, 3, 8, 1, -6}

	fmt.Print("\n")
	mayor, menor := lista.ObtenerMayorYMenor(list)
	fmt.Println("el valor mas grande es: ", mayor, " y el menor es: ", menor)
}
