package Chapter02Sorting

func QuickSort(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	key := a[left]
	for left < right {
		for left < right && a[right] >= key {
			right--
		}
		if left < right {
			a[left] = a[right]
			left++
		}
		for left < right && a[left] <= key {
			left++
		}
		if left < right {
			a[right] = a[left]
			right--
		}
	}
	a[left] = key
	quickSort(a, i, left-1)
	quickSort(a, left+1, j)
}