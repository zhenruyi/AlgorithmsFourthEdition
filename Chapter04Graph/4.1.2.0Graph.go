package Chapter04Graph

import "AlgorithmsFourthEdition/Chapter01Foundation"

type Graph struct {
	V int
	E int
	adj []Chapter01Foundation.ListBag
}

func NewGraph(v int) *Graph {
	 return &Graph{
		V: v,
		E: 0,
		adj: make([]Chapter01Foundation.ListBag, v),
	}
}

func (g *Graph) AddEdge(v, w int) {
	g.adj[v].Add(w)
	g.adj[w].Add(v)
	g.E++
}

func (g *Graph) Adj(v int) []interface{} {
	return g.adj[v].Iterator()
}