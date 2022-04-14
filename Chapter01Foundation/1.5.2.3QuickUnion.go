package Chapter01Foundation

func (uf *UnionFind) QuickUnionFind(p int) int {
	// 找到根触点
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf *UnionFind) QuickUnionUnion(p, q int) {
	pRoot := uf.QuickUnionFind(p)
	qRoot := uf.QuickUnionFind(q)
	if pRoot == qRoot {
		return
	}
	uf.id[pRoot] = qRoot
	uf.count--
}