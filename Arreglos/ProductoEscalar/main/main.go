package main

import (
	"fmt"
	"productoEscalar/productoEscalar"
)

func main() {

	arr1 := []int{1, 2, 3}
	arr2 := []int{2, 4, 6}

	totalSuma, prodEscalar := productoEscalar.CalcularProductoEscalar(arr1, arr2)

	fmt.Println(totalSuma, prodEscalar)
}
