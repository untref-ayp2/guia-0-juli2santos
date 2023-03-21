package main

import (
	"factorial/factorial"
	"fmt"
)

func main() {

	var valor int
	fmt.Scanf("%d", &valor)

	result := factorial.CalculaFactorial(valor)

	fmt.Println(result)

}
