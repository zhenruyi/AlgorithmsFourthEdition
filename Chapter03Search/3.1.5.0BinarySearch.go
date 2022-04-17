package Chapter03Search

type BinarySearchST struct {
	keys []int
	vals []interface{}
	size int
}

func (st *BinarySearchST) Size() int {
	return st.size
}

func (st *BinarySearchST) IsEmpty() bool {
	return st.size == 0
}

func (st *BinarySearchST) Rank(key int) int {
	if st.size == 0 {
		return 0
	}
	lo, hi := 0, st.size-1
	for lo <= hi{
		mid := lo + (hi-lo)/2
		if st.keys[mid] == key {
			return mid
		}
		if key < st.keys[mid] {
			hi = mid-1
		} else {
			lo = mid+1
		}
	}
	return lo
}

func (st *BinarySearchST) rankRecursion(key int, lo, hi int) int {
	if hi < lo {
		return lo
	}
	mid := lo + (hi - lo) / 2
	if key == st.keys[mid] {
		return mid
	}
	if key < st.keys[mid] {
		return st.rankRecursion(key, lo, mid-1)
	} else {
		return st.rankRecursion(key, mid+1, hi)
	}
}

func (st *BinarySearchST) Get(key int) interface{} {
	if st.IsEmpty() {
		return nil
	}
	i := st.Rank(key)
	if i < st.size && st.keys[i] == key {
		return st.vals[i]
	}
	return nil
}

func (st *BinarySearchST) Put(key int, val interface{}) {
	rank := st.Rank(key)
	if st.keys[rank] == key {
		st.vals[rank] = val
	}
	st.keys = append(st.keys, 0)
	st.vals = append(st.vals, nil)
	st.size++
	for i := st.size-1; i > rank; i++ {
		st.keys[i] = st.keys[i-1]
		st.vals[i] = st.vals[i-1]
	}
	st.keys[rank] = key
	st.vals[rank] = val
}

func (st *BinarySearchST) Min() int {
	return st.keys[0]
}

func (st *BinarySearchST) Max() int {
	return st.keys[st.size-1]
}

func (st *BinarySearchST) Select(k int) int {
	return st.keys[k]
}

func (st *BinarySearchST) Ceiling(key int) int {
	rank := st.Rank(key)
	return rank
}

func (st *BinarySearchST) Floor(key int) int {
	rank := st.Rank(key)
	if st.keys[rank] == key {
		return rank
	}
	return rank-1
}

func (st *BinarySearchST) Del(key int) {
	rank := st.Rank(key)
	if st.keys[rank] != key {
		return
	}
	pre := st.keys[:rank]
	suf := st.keys[rank+1:]
	st.size--
	st.keys = append(pre, suf...)
}
