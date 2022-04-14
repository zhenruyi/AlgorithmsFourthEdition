package Chapter01Foundation

type UnionFind struct {
	id []int
	count int  // 分量的数量
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func (uf *UnionFind) Connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func (uf *UnionFind) find(p int) int {
	return 0
}

func (uf *UnionFind) union(p, q int) {

}