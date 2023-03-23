package main

import (
	"fmt"
	"union/union"
)

func main() {
	arreglo1 := []int{1, 2, 3}
	arreglo2 := []int{4, 5, 6}
	unionArr := union.UnionArr(arreglo1, arreglo2)
	interseccionArr := union.InterseccionArr(arreglo1, arreglo2)
	fmt.Println("La union es:", unionArr)
	fmt.Println("La interseccion es:", interseccionArr)
}
