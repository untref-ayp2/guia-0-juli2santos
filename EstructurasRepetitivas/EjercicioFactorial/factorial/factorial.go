package factorial

func CalculaFactorial(valor int) int {
	if valor == 0 {
		return 1
	}

	return valor * CalculaFactorial(valor-1)
}
