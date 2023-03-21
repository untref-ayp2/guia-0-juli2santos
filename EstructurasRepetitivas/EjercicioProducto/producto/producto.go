package producto

func Multiplicar(v1 int, v2 int) int {

	total := 0

	var i int = 0

	// esto va de 0 a v2

	for i < v2 {
		total += v1
		i++
	}

	return total

}
