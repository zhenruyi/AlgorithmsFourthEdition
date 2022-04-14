package Chapter01Foundation

func (uf *UnionFind) QuickFindFind(p int) int {
	return uf.id[p]
}

// 将p和q归并到相同的分量中
func (uf *UnionFind) QuickFindUnion(p, q int) {
	// p和q的分量
	pID := uf.find(p)
	qID := uf.find(q)
	// 相同则返回
	if pID == qID {
		return
	}
	// 将和p分量相同的变成q的分量
	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
	// 分量减少一个
	uf.count--
}
