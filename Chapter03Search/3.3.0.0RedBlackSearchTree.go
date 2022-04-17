package Chapter03Search

const (
	RED = true
	BLACK = false
)

type RBNode struct {
	key int
	val interface{}
	left, right *RBNode
	size int
	color bool
}

func NewRBNode(k int, v interface{}) *RBNode {
	return &RBNode{
		key: k,
		val: k,
		size: 1,
		color: RED,
		left: nil,
		right: nil,
	}
}

type RedBlackT struct {
	root *RBNode
}

func (t *RedBlackT) isRed(h *RBNode) bool {
	return h.color
}

func (t *RedBlackT) rotateLeft(h *RBNode) *RBNode {
	x := h.right
	h.right = x.left
	x.left = h
	x.color = h.color
	h.color = RED
	x.size = h.size
	h.size = h.left.size + h.right.size + 1
	return x
}

func (t *RedBlackT) rotateRight(h *RBNode) *RBNode {
	x := h.left
	h.left = x.right
	x.right = h
	x.color = h.color
	h.color = RED
	x.size = h.size
	h.size = h.left.size + h.right.size + 1
	return x
}

func (t *RedBlackT) flipColors(h *RBNode) {
	h.color = RED
	h.left.color = BLACK
	h.right.color = BLACK
}

func (t *RedBlackT) Put(key int, val interface{}) {
	t.root = t.put(t.root, key, val)
	t.root.color = BLACK
}

func (t *RedBlackT) put(h *RBNode, key int, val interface{}) *RBNode {
	if h == nil {
		return NewRBNode(key, val)
	}
	if key == h.key {
		h.val = val
	}
	if key < h.key {
		h.left = t.put(h.left, key, val)
	}
	if key > h.key {
		h.right = t.put(h.right, key, val)
	}
	if h.right.color && !h.left.color {
		h = t.rotateLeft(h)
	}
	if h.left.color && h.left.left.color {
		h = t.rotateRight(h)
	}
	if h.left.color && h.right.color {
		t.flipColors(h)
	}
	h.size = h.left.size + h.right.size + 1
	return h
}

