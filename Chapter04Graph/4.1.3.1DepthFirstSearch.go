package Chapter04Graph

type DepthFirstSearch struct {
	marked []bool
	count int
}

func NewDepthFirstSearch(g Graph, s int) *DepthFirstSearch {
	d := new(DepthFirstSearch)
	d.count = 0
	d.marked = make([]bool, g.V)
	d.DFS(g, s)
	return d
}

func (d *DepthFirstSearch) DFS(g Graph, v int) {
	d.marked[v] = true
	d.count++
	for _, w := range g.Adj(v) {
		if !d.marked[w.(int)] {
			d.DFS(g, w.(int))
		}
	}
}