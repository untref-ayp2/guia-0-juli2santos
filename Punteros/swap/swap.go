package swap

func Swap(px, py *int) {
	aux := *px
	*px = *py
	*py = aux

}
