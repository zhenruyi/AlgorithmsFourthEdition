package Chapter01Foundation

type Node struct {
	data interface{}
	next *Node
}

type ListStack struct {
	first *Node
	length int
}

func (l *ListStack) IsEmpty() bool {
	return l.length == 0
}

func (l *ListStack) Size() int {
	return l.length
}

func (l *ListStack) Push(elem interface{}) {
	node := new(Node)
	node.data = elem
	node.next = l.first
	l.first = node
	l.length++
}

func (l *ListStack) Pop() interface{} {
	if l.IsEmpty() {
		return nil
	}
	ret := l.first
	l.first = l.first.next
	l.length--
	return ret.data
}
