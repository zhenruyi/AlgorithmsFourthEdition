package Chapter02Sorting


func less(a []int, i, j int) bool {
	return a[i] < a[j]
}

func exch(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

