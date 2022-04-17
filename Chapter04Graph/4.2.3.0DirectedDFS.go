package Chapter04Graph

type DirectedDFS struct {
	marked []bool
}

func NewDirectedDFS(g Graph, s int) *DirectedDFS {
	d := new(DirectedDFS)
	d.marked = make([]bool, g.V)
	d.DFS(g, s)
	return d
}

func (d *DirectedDFS) DFS(g Graph, v int) {
	d.marked[v] = true
	for _, w := range g.Adj(v) {
		if !d.marked[w.(int)] {
			d.DFS(g, w.(int))
		}
	}
}