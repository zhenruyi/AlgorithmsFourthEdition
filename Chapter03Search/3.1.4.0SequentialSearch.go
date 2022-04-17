package Chapter03Search

type Node struct {
	key interface{}
	val interface{}
	nxt *Node
}

func NewNode(k, v interface{}, n *Node) *Node {
	return &Node{
		key: k,
		val: v,
		nxt: n,
	}
}

type SequentialSearchST struct {
	first *Node
}

func (s *SequentialSearchST) Get(key interface{}) interface{} {
	for x := s.first; x != nil; x = x.nxt {
		if x.key == key {
			return x.val
		}
	}
	return nil
}

func (s *SequentialSearchST) Put(key, val interface{}) {
	for x := s.first; x != nil; x = x.nxt {
		if x.key == key {
			x.val = val
			return
		}
	}
	s.first = NewNode(key, val, s.first)
}

