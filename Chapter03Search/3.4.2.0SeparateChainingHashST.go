package Chapter03Search

type ChainingHT struct {
	// 链表的平均长度N/M
	// 键值对总数
	N int
	// 数组的大小
	M int
	// 数组
	st []SequentialSearchST
}

func NewChainingHT(m int) *ChainingHT {
	ret := &ChainingHT{
		M: m,
		st: make([]SequentialSearchST, m),
	}
	return ret
}

func (ht *ChainingHT) hashfunc(key int) int {
	return key % ht.M
}

func (ht *ChainingHT) Get(key int) interface{} {
	return ht.st[ht.hashfunc(key)].Get(key)
}

func (ht *ChainingHT) Put(key int, val interface{}) {
	ht.st[ht.hashfunc(key)].Put(key, val)
}