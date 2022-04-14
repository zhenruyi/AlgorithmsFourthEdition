package Chapter02Sorting

func SelectionSort(a []int) {
	length := len(a)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i+1; j < length; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		exch(a, min, i)
	}
}

func InsertionSort(a []int) {
	length := len(a)
	for i := 1; i < length; i++ {
		for j := i; j >= 0; j-- {
			if a[j] < a[j-1] {
				exch(a, j, j-1)
			}
		}
	}
}

func ShellSort(a []int) {
	length := len(a)
	h := 1
	for h < length / 3 {
		h = h * 3 + 1
	}
	for h >= 1 {
		for i := h; i < length; i++ {
			for j := i; j >= h; j -= h {
				if less(a, j, j-h) {
					exch(a, j, j-h)
				}
			}
		}
		h /= 3
	}
}