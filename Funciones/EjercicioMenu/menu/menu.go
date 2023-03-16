package menu

import "fmt"

func GenerarMenu(opcion int) {

	switch opcion {
	case 1:
		fmt.Println("el valor seleccionado es:", opcion)
	case 2:
		fmt.Println("el valor seleccionado es:", opcion)
	case 3:
		fmt.Println("el valor seleccionado es:", opcion)
	case 4:
		fmt.Println("el valor seleccionado es:", opcion)
	default:
		fmt.Println("Error, el numero no esta en las opciones disponibles")

	}
}
