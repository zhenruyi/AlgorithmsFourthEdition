package Chapter04Graph

import (
	"AlgorithmsFourthEdition/Chapter01Foundation"
	"container/list"
)

type LazyPrimMST struct {
	marked []bool	// 顶点
	mst *Chapter01Foundation.ListQueue  // 边
	// 横切边，原文是最小队列，我是最大队列
	//pq *Chapter02Sorting.MaxPriorityQueue
	pq *list.List
}

func (p *LazyPrimMST) Visit(g EdgeWeightedGraph, v int) {
	p.marked[v] = true
	for _, row := range g.Adj(v) {
		e := row.(Edge)
		if !p.marked[e.Other(v)] {
			p.pq.PushFront(e)
			p.SortList()
		}
	}
}

func (p *LazyPrimMST) SortList() {
	if p.pq.Len() == 0 {
		return
	}
	min := p.pq.Front()
	for po := p.pq.Front(); po != nil; po = po.Next() {
		if po.Value.(Edge).weight < min.Value.(Edge).weight {
			min = po
		}
	}
	p.pq.MoveBefore(min, p.pq.Front())
}

func NewLazyPrimMST(g EdgeWeightedGraph) *LazyPrimMST {
	ret := new(LazyPrimMST)
	ret.pq = list.New()
	ret.marked = make([]bool, g.V)
	ret.mst = new(Chapter01Foundation.ListQueue)
	ret.Visit(g, 0)
	for ret.pq.Len() != 0 {
		e := ret.pq.Remove(ret.pq.Front()).(Edge)
		ret.SortList()
		v := e.Either()
		w := e.Other(v)
		if ret.marked[v] && ret.marked[w] {
			continue
		}
		ret.mst.Enqueue(e)
		if !ret.marked[v] {
			ret.Visit(g, v)
		}
		if !ret.marked[w] {
			ret.Visit(g, w)
		}
	}
	return ret
}