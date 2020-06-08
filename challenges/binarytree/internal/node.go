package internal

import "fmt"

// A node is a structure that holds data and relation information for it's children.
type Node struct {
	data  int64
	left  *Node
	right *Node
}

// a tree is a binary node structure with a root.
type BinaryTree struct {
	root *Node
}

// return the root of the binary tree
func NewBinaryTree() BinaryTree {
	return BinaryTree{root: nil}
}

// dump the full tree
func (t *BinaryTree) String() string {
	format := "---\n"
	//format += fmt.Sprintf("ROOT %d\n", t.root.data)
	format += t.root.Stringify(0)
	format += "---\n"
	return format
}

func (n *Node) Stringify(level int64) string {
	format := ""
	if n != nil {
		for i := int64(0); i < level; i++ {
			format += "     "
		}
		format += fmt.Sprintf(format+"(%d)---[ %s \n", level, n.String())
		level++
		format += n.left.Stringify(level)
		format += n.right.Stringify(level)
	}
	return format
}

func (n *Node) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("Node: %d", n.data)
}

func (t *BinaryTree) Search(f int64) bool {

	return t.root.Search(f)

}

// search the tree for the value
func (n *Node) Search(v int64) bool {
	if n == nil {
		return false
	}

	if v < n.data {
		return n.left.Search(v)
	}

	if v > n.data {
		return n.right.Search(v)
	}

	return true
}

func (t *BinaryTree) Insert(data int64) *BinaryTree {

	if t.root == nil {
		t.root = &Node{data: data}
		return t
	}

	t.root.Insert(data)
	return t
}

func (n *Node) Insert(data int64) {

	if n == nil {
		return
	}

	if data <= n.data {
		if n.left == nil {
			n.left = &Node{data: data, left: nil, right: nil}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &Node{data: data, left: nil, right: nil}
		} else {
			n.right.Insert(data)
		}
	}
}
