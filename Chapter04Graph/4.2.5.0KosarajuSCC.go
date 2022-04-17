package Chapter04Graph

type KasarajuSCC struct {
	marked []bool
	id []int	//连通分量的编号
	count int	//当前编号
}

func NewKasarajuSCC(g Graph) *KasarajuSCC {
	ret := new(KasarajuSCC)
	ret.marked = make([]bool, g.V)
	ret.count = 0
	ret.id = make([]int, g.V)
	/*
	for s := 0; s < g.V; s++ {
		if !ret.marked[s] {
			ret.DFS(g, s)
			ret.count++
		}
	}
	*/
	return ret
}

func (cc *KasarajuSCC) DFS(g Graph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	for _, wRow := range g.Adj(v) {
		w := wRow.(int)
		if !cc.marked[w] {
			cc.DFS(g, w)
		}
	}
}

func (cc *KasarajuSCC) StronglyConnected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}