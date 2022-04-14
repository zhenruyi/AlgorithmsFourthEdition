package Chapter01Foundation

type ListBag struct {
	first *Node
}

func (l *ListBag) Add(item interface{}) {
	node := new(Node)
	node.data = item
	node.next = l.first
	l.first = node
}

// 迭代，返回一个切片
func (l *ListBag) Iterator() []interface{} {
	ret := make([]interface{}, 0)
	for node := l.first; node != nil; node = node.next {
		ret = append(ret, node.data)
	}
	return ret
}