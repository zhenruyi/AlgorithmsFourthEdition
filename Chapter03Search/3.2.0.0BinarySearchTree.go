package Chapter03Search

type TreeNode struct {
	key int
	val interface{}
	left, right *TreeNode
	size int
}

func NewTreeNode(k int, v interface{}) *TreeNode {
	return &TreeNode{
		key: k,
		val: v,
		size: 1,
		left: nil,
		right: nil,
	}
}

func (t *TreeNode) Size() int {
	return t.size
}

type BinarySearchTree struct {
	root *TreeNode
}

func (t BinarySearchTree) Size() int {
	return t.root.size
}

func (t *BinarySearchTree) Get(key int) interface{} {
	return t.get(t.root, key)
}

func (t *BinarySearchTree) get(root *TreeNode, key int) interface{} {
	if root == nil {
		return nil
	}
	if key == root.key {
		return root.val
	}
	if key < root.key {
		return t.get(root.left, key)
	} else {
		return t.get(root.right, key)
	}
}

func (t *BinarySearchTree) Put(key int, val interface{}) {
	t.root = t.put(t.root, key, val)
}

func (t *BinarySearchTree) put(root *TreeNode, key int, val interface{}) *TreeNode {
	if root == nil {
		return NewTreeNode(key, val)
	}
	if key == root.key {
		root.val = val
	}
	if key < root.key {
		root.left = t.put(root.left, key, val)
	} else {
		root.right = t.put(root.right, key, val)
	}
	root.size = root.left.size + root.right.size + 1
	return root
}

func (t *BinarySearchTree) Min() int {
	return t.min(t.root)
}

func (t *BinarySearchTree) min(root *TreeNode) int {
	if root.left == nil {
		return root.key
	}
	return t.min(root.left)
}

func (t *BinarySearchTree) Max() int {
	return t.max(t.root)
}

func (t *BinarySearchTree) max(root *TreeNode) int {
	if root.right == nil {
		return root.key
	}
	return t.max(root.right)
}

func (t *BinarySearchTree) Floor(key int) int {
	return t.floor(t.root, key)
}

func (t *BinarySearchTree) floor(root *TreeNode, key int) int {
	if root == nil {
		return -1
	}
	if key == root.key {
		return root.key
	}
	if key < root.key {
		return t.floor(root.left, key)
	}
	if root.right == nil {
		return root.key
	}
	if rightKey := root.right.key; rightKey > key {
		return root.key
	}
	return t.floor(root.right, key)
}

func (t *BinarySearchTree) Ceiling(key int) int {
	return t.ceiling(t.root, key)
}

func (t *BinarySearchTree) ceiling(root *TreeNode, key int) int {
	if root == nil {
		return -1
	}
	if key == root.key {
		 return key
	}
	if key > root.key {
		return t.floor(root.right, key)
	}
	if root.left == nil {
		return root.key
	}
	if leftKey := root.left.key; leftKey < key {
		return root.key
	}
	return t.floor(root.left, key)
}

func (t *BinarySearchTree) Sel(rank int) int {
	return t.sel(t.root, rank)
}

func (t *BinarySearchTree) sel(root *TreeNode, rank int) int {
	if root == nil {
		return -1
	}
	l := root.left.size
	if l == rank {
		return root.key
	}
	if l > rank {
		return t.sel(root.left, rank)
	}
	return t.sel(root.right, rank-l-1)
}

func (t *BinarySearchTree) DelMax() {
	t.root = t.delMax(t.root)
}

func (t *BinarySearchTree) delMax(root *TreeNode) *TreeNode {
	if root.right == nil {
		return root.left
	}
	root.right = t.delMax(root.right)
	root.size = root.left.size + root.right.size + 1
	return root
}

func (t *BinarySearchTree) Del(key int) {
	t.root = t.del(t.root, key)
}

func (t *BinarySearchTree) del(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if key < root.key {
		root.left = t.del(root.left, key)
	}
	if key > root.key {
		root.right = t.del(root.right, key)
	}
	if root.right == nil {
		return root.left
	}
	if root.left == nil {
		return root.right
	}
	// 删除现在的结点
	// 右子结点的结点作为现在的结点
	// 处理size
}