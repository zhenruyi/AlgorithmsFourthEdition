package Chapter04Graph

type Edge struct {
	v, w int
	weight float64
}

func NewEdge(v, w int, wgh float64) *Edge {
	return &Edge{
		v: v,
		w: w,
		weight: wgh,
	}
}

func (e *Edge) Either() int {
	return e.v
}

func (e *Edge) Other(v int) int {
	if v == e.v {
		return e.w
	}
	if v == e.w {
		return e.v
	}
	return -1
}

func (e *Edge) CmpTo(that Edge) int {
	if e.weight > that.weight {
		return 1
	}
	if e.weight < that.weight {
		return -1
	}
	return 0
}