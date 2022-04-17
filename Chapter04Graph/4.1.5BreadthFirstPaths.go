package Chapter04Graph

import "container/list"

type BreadthFirstPaths struct {
	marked []bool
	edgeTo []int
	start int
}

func NewBreadthFirstPaths(g Graph, s int) *BreadthFirstPaths {
	ret := new(BreadthFirstPaths)
	ret.marked = make([]bool, g.V)
	ret.edgeTo = make([]int, g.V)
	ret.start = s
	ret.BFS(g, s)
	return ret
}

func (p *BreadthFirstPaths) BFS(g Graph, s int) {
	l := list.New()
	p.marked[s] = true
	l.PushBack(s)
	for l.Len() != 0 {
		v := l.Remove(l.Front()).(int)
		for _, wRow := range g.Adj(v) {
			w := wRow.(int)
			if !p.marked[w] {
				p.edgeTo[w] = v
				p.marked[w] = true
				l.PushBack(w)
			}
		}
	}
}

func (p *BreadthFirstPaths) HasPathTo(v int) bool {
	return p.marked[v]
}

func (p *BreadthFirstPaths) PathTo(v int) []int {
	if !p.HasPathTo(v) {
		return nil
	}
	ret := make([]int, 0)
	for x := v; x != p.start; x = p.edgeTo[x] {
		ret = append(ret, x)
	}
	ret = append(ret, p.start)
	return ret
}