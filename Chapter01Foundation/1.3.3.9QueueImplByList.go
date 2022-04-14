package Chapter01Foundation

type ListQueue struct {
	first *Node
	last *Node
	length int
}

func (l *ListQueue) IsEmpty() bool {
	return l.length == 0
}

func (l *ListQueue) Size() int {
	return l.length
}

func (l *ListQueue) Enqueue(item interface{}) {
	node := new(Node)
	node.data = item
	node.next = nil
	if l.length == 0 {
		l.first = node
		l.last = node
	} else {
		l.last.next = node
		l.last = node
	}
	l.length++
}

func (l *ListQueue) Dequeue() interface{} {
	if l.IsEmpty() {
		return nil
	}
	ret := l.first.data
	l.first = l.first.next
	l.length--
	if l.IsEmpty() {
		l.last = nil
	}
	return ret
}
