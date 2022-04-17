package Chapter04Graph

import "container/list"

type DirectedCycle struct {
	marked []bool
	edgeTo []int
	cycle *list.List
	onStack []bool
}

func NewDirectedCycle(g Graph) *DirectedCycle {
	ret := &DirectedCycle{
		onStack: make([]bool, g.V),
		edgeTo: make([]int, g.V),
		marked: make([]bool, g.V),
	}
	for v:=0; v < g.V; v++ {
		if !ret.marked[v] {
			ret.DFS(g, v)
		}
	}
	return ret
}

func (d *DirectedCycle) DFS(g Graph, v int) {
	d.onStack[v] = true
	d.marked[v] = true
	for _, row := range g.Adj(v) {
		w := row.(int)
		if d.HasCycle() {
			return
		} else if !d.marked[w] {
			d.edgeTo[w] = v
			d.DFS(g, w)
		} else if d.onStack[w] {
			d.cycle = list.New()
			for x:=v; x != w; x = d.edgeTo[x] {
				d.cycle.PushFront(x)
			}
			d.cycle.PushFront(w)
			d.cycle.PushFront(v)
		}
	}
}

func (d *DirectedCycle) HasCycle() bool {
	return d.cycle.Len() != 0
}

func (d *DirectedCycle) Cycle() *list.List {
	return d.cycle
}