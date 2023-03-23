package suma

func SumArray(arr []int) int {
	var total int = 0
	for _, num := range arr {
		total += num
	}
	return total
}
