package main

import (
	"fmt"
	"swap/swap"
)

func main() {

	fmt.Printf("")
	valor1 := 123
	valor2 := 234

	v1 := &valor1
	v2 := &valor2
	swap.Swap(v1, v2)

}
