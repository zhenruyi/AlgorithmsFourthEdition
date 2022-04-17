package Chapter04Graph

// 是否有一条路径
type DepthFirstPaths struct {
	marked []bool	//这个顶点上调用过DFS吗？
	edgeTo []int	//从起点到一个顶点的已知路径上的最后一个顶点
	s int			//起点
}

func NewDepthFirstPaths(g Graph, start int) *DepthFirstPaths {
	p := new(DepthFirstPaths)
	p.marked = make([]bool, g.V)
	p.edgeTo = make([]int, g.V)
	p.s = start
	p.DFS(g, start)
	return p
}

func (p *DepthFirstPaths) DFS(g Graph, v int) {
	p.marked[v] = true
	for _, w := range g.Adj(v) {
		if !p.marked[w.(int)] {
			p.edgeTo[w.(int)] = v
			p.DFS(g, w.(int))
		}
	}
}

func (p *DepthFirstPaths) HasPathTo(v int) bool {
	return p.marked[v]
}

func (p *DepthFirstPaths) PathTo(v int) []int {
	if !p.HasPathTo(v) {
		return nil
	}
	ret := make([]int, 0)
	for x := v; x != p.s; x = p.edgeTo[x] {
		ret = append(ret, x)
	}
	ret = append(ret, p.s)
	return ret
}
