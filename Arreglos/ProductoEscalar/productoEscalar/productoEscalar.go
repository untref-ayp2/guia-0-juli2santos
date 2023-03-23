package productoEscalar

func CalcularProductoEscalar(arr1 []int, arr2 []int) (int, int) {
	var sum int = 0
	var productoE int = 0
	for i := 0; i < len(arr1); i++ {
		sum += arr1[i] + arr2[i]
		productoE += arr1[i] * arr2[i]
	}
	return sum, productoE
}
