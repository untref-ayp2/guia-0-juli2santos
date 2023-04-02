package ejercicio

import (
	figuras "Figuras/Figuras"
	punto "Figuras/Punto"
	rectangulo "Figuras/Rectangulo"
	"Figuras/circulo"
	"Figuras/cuadrado"
	"fmt"
	"strings"
)

func Ejercicio() {

	var (
		opcion      string
		arrOpciones []string
		arrFiguras  = [5]figuras.Figura{}
	)

	fmt.Println(" 1- Rectangulo\n	2- Cuadrado\n	3- Circulo")

	for i := 0; i < 5; i++ {

		fmt.Scanln(&opcion)
		opcion = strings.ToUpper(opcion)
		fmt.Println("Usted a seleccionado la opcion:", opcion)

		switch {
		case opcion == "A":
			arrOpciones = append(arrOpciones, opcion)
		case opcion == "B":
			arrOpciones = append(arrOpciones, opcion)
		case opcion == "C":
			arrOpciones = append(arrOpciones, opcion)
		default:
			i--
			fmt.Println("valor invalido")

		}
	}

	for i := 0; i < 5; i++ {
		switch {
		case arrOpciones[i] == "A":
			arrFiguras[i] = crearRectangulo(i)
		case arrOpciones[i] == "B":
			arrFiguras[i] = crearCuadrado(i)
		case arrOpciones[i] == "C":
			arrFiguras[i] = crearCirculo(i)
		}
	}

	for _, figura := range arrFiguras {
		mostrar(figura)
	}
	fmt.Println("")
}

func crearCirculo(i int) figuras.Figura {
	fmt.Println("Ingrese el punto inicial ")
	var puntoScan int
	var puntoVal punto.Punto
	fmt.Scan(&puntoScan)
	puntoVal.X = float32(puntoScan)
	fmt.Scan(&puntoScan)
	puntoVal.Y = float32(puntoScan)
	fmt.Println("Ingrese el radio ")
	fmt.Scan(&puntoScan)
	return circulo.Circulo{P: puntoVal, R: float32(puntoScan)}
}

func crearCuadrado(i int) figuras.Figura {
	fmt.Println("Ingrese los parametros de la figura ", (i + 1), "[CUADRADO]\nPunto: ")
	var puntoScan int
	var puntoVal punto.Punto
	fmt.Scan(&puntoScan)
	puntoVal.X = float32(puntoScan)
	fmt.Scan(&puntoScan)
	puntoVal.Y = float32(puntoScan)
	fmt.Println("Ingrese los parametros de la figura ", (i + 1), "[CUADRADO]\nTamaño de lado: ")
	fmt.Scan(&puntoScan)
	return cuadrado.Cuadrado{Pto: puntoVal, Lado: float32(puntoScan)}
}

func crearRectangulo(i int) figuras.Figura {
	fmt.Println("Ingrese los parametros de la figura", (i + 1), "[RECTANGULO]\nPunto 1: ")
	var puntoScan int
	var puntoVal punto.Punto
	fmt.Scan(&puntoScan)
	puntoVal.X = float32(puntoScan)
	fmt.Scan(&puntoScan)
	puntoVal.Y = float32(puntoScan)
	fmt.Println("Ingrese los parametros de la figura", (i + 1), "[RECTANGULO]\nPunto 2: ")
	var puntoDos punto.Punto
	fmt.Scan(&puntoScan)
	puntoDos.X = float32(puntoScan)
	fmt.Scan(&puntoScan)
	puntoDos.Y = float32(puntoScan)
	return rectangulo.Rectangulo{P1: puntoVal, P2: puntoDos}
}

func mostrar(f figuras.Figura) {
	fmt.Println(f.ToString())
	fmt.Println("Area: ", f.Area())
	fmt.Println("Perimetro: ", f.Perimetro())
	fmt.Println()
}
