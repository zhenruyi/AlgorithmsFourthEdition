package Chapter02Sorting

// 归并a和b两个切片
func merge1(a, b []int) []int {
	ret := make([]int, 0)
	for len(a) > 0 && len(b) > 0 {
		if a[0] > b[0] {
			ret = append(ret, b[0])
			b = b[1:]
		} else {
			ret = append(ret, a[0])
			a = a[1:]
		}
	}
	for len(a) > 0 {
		ret = append(ret, a[0])
		a = a[1:]
	}
	for len(b) > 0 {
		ret = append(ret, b[0])
		b = b[1:]
	}
	return ret
}

// 归并[low:mid]和[mid+1:hig]
func merge2(a []int, low, mid, hig int) {
	i, j := low, mid+1
	aux := make([]int, len(a))
	for index, v := range a {
		aux[index] = v
	}
	for k := low; k <= hig; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hig {
			a[k] = aux[i]
			i++
		} else if less(aux, i, j) {
			a[k] = aux[i]
			i++
		} else {
			a[k] = aux[j]
			j++
		}
	}
}

func MergeSortTop2Bottom(a []int) []int {
	length := len(a)
	if length <= 1 {
		return a
	}
	mid := length / 2
	return merge1(MergeSortTop2Bottom(a[:mid]), MergeSortTop2Bottom(a[mid:]))
}

func MergeSortBottom2Top(a []int) {
	length := len(a)
	for sz := 1; sz < length; sz += sz {
		for low := 0; low < length-sz; low += sz+sz {
			merge2(a, low, low+sz-1, min(low+sz+sz-1, length-1))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}