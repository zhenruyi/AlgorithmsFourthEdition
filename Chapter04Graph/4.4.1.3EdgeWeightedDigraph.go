package Chapter04Graph

import "container/list"

type EdgeWeightedDigraph struct {
	V int
	E int
	adj []list.List
}

func NewEdgeWeightedDigraph(v int) *EdgeWeightedDigraph {
	return &EdgeWeightedDigraph{v, 0, make([]list.List, v)}
}

func (dg *EdgeWeightedDigraph) AddEdge(e DirectedEdge) {
	dg.adj[e.From()].PushFront(e)
	dg.E++
}

func (dg *EdgeWeightedDigraph) Adj(v int) *list.List {
	return &dg.adj[v]
}

func (dg *EdgeWeightedDigraph) Edges() *list.List {
	ret := list.New()
	for v := 0; v < dg.V; v++ {
		l := dg.adj[v]
		for tmp := l.Front(); tmp != nil; tmp = tmp.Next() {
			ret.PushBack(tmp.Value)
		}
	}
	return ret
}
