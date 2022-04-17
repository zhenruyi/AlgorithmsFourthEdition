package Chapter04Graph

import "AlgorithmsFourthEdition/Chapter01Foundation"

type EdgeWeightedGraph struct {
	V int
	E int
	adj []Chapter01Foundation.ListBag
}

func NewEdgeWeightedGraph(v int) *EdgeWeightedGraph {
	return &EdgeWeightedGraph{
		V: v,
		E: 0,
		adj: make([]Chapter01Foundation.ListBag, v),
	}
}

func (g *EdgeWeightedGraph) AddEdge(e Edge) {
	v, w := e.v, e.w
	g.adj[v].Add(e)
	g.adj[w].Add(e)
	g.E++
}

func (g *EdgeWeightedGraph) Adj(v int) []interface{} {
	ret := g.adj[v].Iterator()
	return ret
}