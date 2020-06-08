package internal

import (
	"encoding/json"
	"fmt"
)

// A node is a structure that holds data and relation information for it's children.
type Node struct {
	Data  int64 `json:"data,omitempty"`
	Left  *Node `json:"left,omitempty"`
	Right *Node `json:"right,omitempty"`
}

// a tree is a binary node structure with a root.
type BinaryTree struct {
	Root *Node `json:"root,omitempty"`
}

// return the root of the binary tree
func NewBinaryTree() BinaryTree {
	return BinaryTree{Root: nil}
}

// dump the full tree
func (t *BinaryTree) String() string {
	format := "---\n"
	//format += fmt.Sprintf("ROOT %d\n", t.root.data)
	format += t.Root.Stringify(0)
	format += "---\n"
	return format
}

func (t *BinaryTree) Serialize() string {
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func Deserialize(s string) *BinaryTree {

	fmt.Println(s)

	t := &BinaryTree{}
	if err := json.Unmarshal([]byte(s), t); err != nil {
		panic(err)
	}
	return t
}

// Print out the tree
func (n *Node) Stringify(level int64) string {
	format := ""
	if n != nil {
		for i := int64(0); i < level; i++ {
			format += "     "
		}
		format += fmt.Sprintf(format+"(%d)---[ %s \n", level, n.String())
		level++
		format += n.Left.Stringify(level)
		format += n.Right.Stringify(level)
	}
	return format
}

// Implement the interface to print the object
func (n *Node) String() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("Node: %d", n.Data)
}

// search the nodes for an item
func (t *BinaryTree) Search(f int64) bool {

	return t.Root.Search(f)

}

// search the tree for the value
func (n *Node) Search(v int64) bool {
	if n == nil {
		return false
	}

	if v < n.Data {
		return n.Left.Search(v)
	}

	if v > n.Data {
		return n.Right.Search(v)
	}

	return true
}

// add the data to the tree
func (t *BinaryTree) Insert(data int64) *BinaryTree {

	if t.Root == nil {
		t.Root = &Node{Data: data}
		return t
	}

	t.Root.Insert(data)
	return t
}

// add data to the node
func (n *Node) Insert(data int64) {

	if n == nil {
		return
	}

	if data <= n.Data {
		if n.Left == nil {
			n.Left = &Node{Data: data, Left: nil, Right: nil}
		} else {
			n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Data: data, Left: nil, Right: nil}
		} else {
			n.Right.Insert(data)
		}
	}
}
