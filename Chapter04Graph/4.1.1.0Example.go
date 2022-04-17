package Chapter04Graph

import "fmt"

func Degree(g Graph, v int) int {
	return len(g.Adj(v))
}

func MaxDegree(g Graph) int {
	max := 0
	for i :=0; i < g.V; i++ {
		if d := Degree(g, i); d > max {
			max = d
		}
	}
	return max
}

func AvgDegree(graph Graph) float64 {
	return 2.0 * float64(graph.E) / float64(graph.V)
}

func NOfSelfLoops(graph Graph) int {
	count := 0
	for v := 0; v < graph.V; v++ {
		for _, w := range graph.Adj(v) {
			if w == v {
				count++
			}
		}
	}
	return count / 2
}

func ToString(g Graph) string {
	s := fmt.Sprintf("%v vertices, %v edges\n", g.V, g.E)
	for v := 0; v < g.V; v++ {
		s += fmt.Sprintf("%v: ", v)
		for _, w := range g.Adj(v) {
			s += fmt.Sprintf("%v ", w)
		}
		s += "\n"
	}
	return s
}