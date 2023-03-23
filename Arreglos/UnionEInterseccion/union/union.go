package union

func UnionArr(arrA []int, arrB []int) []int {

	union := arrA

	union = append(union, arrB...)

	return union
}

func InterseccionArr(arrA, arrB []int) []int {

	var interseccion []int

	for _, valorA := range arrA {
		for _, valorB := range arrB {

			if valorA == valorB {
				interseccion = append(interseccion, valorA)
			}
		}
	}

	return interseccion
}
