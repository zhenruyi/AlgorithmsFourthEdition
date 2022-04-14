package Chapter01Foundation

type WeightedQuickUnionUF struct {
	id []int
	size []int
	count int
}

func (uf *WeightedQuickUnionUF) Count() int {
	return uf.count
}

func (uf *WeightedQuickUnionUF) connected(p, q int) bool {
	return uf.find(q) == uf.find(q)
}

func (uf *WeightedQuickUnionUF) find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf *WeightedQuickUnionUF) union(p, q int) {
	pRoot := uf.find(p)
	qRoot := uf.find(q)
	if pRoot == qRoot {
		return
	}
	if uf.size[qRoot] < uf.size[pRoot] {
		uf.id[qRoot] = pRoot
		uf.size[pRoot] += uf.size[qRoot]
	} else {
		uf.id[pRoot] = qRoot
		uf.size[qRoot] += uf.size[pRoot]
	}
	uf.count--
}