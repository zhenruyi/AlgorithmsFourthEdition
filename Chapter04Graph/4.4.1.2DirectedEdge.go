package Chapter04Graph

type DirectedEdge struct {
	v int
	w int
	weight float64
}

func NewDirectedEdge(v, w int, weight float64) *DirectedEdge {
	return &DirectedEdge{v, w, weight}
}

func (E *DirectedEdge) From() int {
	return E.v
}

func (E *DirectedEdge) To() int {
	return E.w
}
