package Chapter04Graph

import "container/list"

type KruskalMST struct {
	mstQueue *list.List
}

// 如果权重最小的边符合要求，则加入到生成树中。
func NewKruskalMST(g EdgeWeightedGraph) *KruskalMST {
	ret := new(KruskalMST)
	return ret
}