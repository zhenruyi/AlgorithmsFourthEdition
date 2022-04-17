package Chapter03Search

type LinearHT struct {
	N int   //键数量
	M int   //数组大小
	keys []int
	vals []interface{}
}

func NewLinearHT(m int) *LinearHT {
	ret := &LinearHT{
		N: 0,
		M: m,
		keys: make([]int, m),
		vals: make([]interface{}, m),
	}
	for i := 0; i < m; i++ {
		ret.keys[i] = -1
		ret.vals[i] = -1
	}
	return ret
}

func (ht *LinearHT) hashFunc(key int) int {
	return key % ht.M
}

func (ht *LinearHT) resize(r int) {
	t := NewLinearHT(r)
	for i := 0; i < ht.M; i++ {
		if ht.keys[i] != -1 {
			t.Put(ht.keys[i], ht.vals[i])
		}
	}
	ht.keys = t.keys
	ht.vals = t.vals
	ht.M = t.M
}

func (ht *LinearHT) Put(key int, val interface{}) {
	if ht.N >= ht.M / 2 {
		ht.resize(2 * ht.M)
	}
	index := ht.hashFunc(key)
	for ht.keys[index] != -1 {
		if ht.keys[index] == key {
			ht.vals[index] = val
			return
		}
		index = (index + 1) % ht.M
	}
	ht.keys[index] = key
	ht.vals[index] = val
	ht.N++
}

func (ht *LinearHT) Get(key int) interface{} {
	for i := ht.hashFunc(key); ht.keys[i] != -1; i = (i+1) % ht.M {
		if ht.keys[i] == key {
			return ht.vals[i]
		}
	}
	return nil
}

// 删除操作
func (ht *LinearHT) Del(key int) {

}