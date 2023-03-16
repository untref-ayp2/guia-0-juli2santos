package main

import (
	"fmt"
	"menu/menu"
)

func main() {
	var valor int
	fmt.Scanf("%d", &valor) // scanf guarda el dato en la direccion de memoria a la q apunta valor
	menu.GenerarMenu(valor)
}
