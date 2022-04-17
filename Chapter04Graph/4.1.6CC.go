package Chapter04Graph

type CC struct {
	marked []bool
	id []int	//连通分量的编号
	count int	//当前编号
}

func NewCC(g Graph) *CC {
	ret := new(CC)
	ret.marked = make([]bool, g.V)
	ret.count = 0
	ret.id = make([]int, g.V)
	for s := 0; s < g.V; s++ {
		if !ret.marked[s] {
			ret.DFS(g, s)
			ret.count++
		}
	}
	return ret
}

func (cc *CC) DFS(g Graph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	for _, wRow := range g.Adj(v) {
		w := wRow.(int)
		if !cc.marked[w] {
			cc.DFS(g, w)
		}
	}
}

func (cc *CC) Connected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}
