package Chapter04Graph

import "AlgorithmsFourthEdition/Chapter01Foundation"

type Digraph struct {
	V int
	E int
	adj []Chapter01Foundation.ListBag
}

func NewDigraph(v int) *Digraph {
	return &Digraph{
		V: v,
		E: 0,
		adj: make([]Chapter01Foundation.ListBag, v),
	}
}

func (g *Digraph) AddEdge(v, w int) {
	g.adj[v].Add(w)
	g.E++
}

func (g *Digraph) Adj(v int) []interface{} {
	return g.adj[v].Iterator()
}

func (g *Digraph) Reverse() *Digraph {
	ret := NewDigraph(g.V)
	for v := 0; v < g.V; v++ {
		for _, row := range g.Adj(v) {
			w := row.(int)
			ret.AddEdge(w, v)
		}
	}
	return ret
}