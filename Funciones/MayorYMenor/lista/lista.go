package lista

func ObtenerMayorYMenor(lista []int) (int, int) {

	mayor := lista[0]
	menor := lista[0]

	for _, valor := range lista {

		if valor > mayor {
			mayor = valor
		}
		if valor < menor {
			menor = valor
		}
	}

	return mayor, menor
}
