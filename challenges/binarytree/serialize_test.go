package binarytree

import (
	"testing"

	"github.com/dathan/go-lang-learning/challenges/binarytree/internal"
)

func TestBinaryTree(t *testing.T) {

	tree := internal.NewBinaryTree()

	type args struct {
		a *internal.BinaryTree
		k int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"tree", args{a: tree.Insert(int64(8)), k: 8}, true},
		{"tree", args{a: tree.Insert(int64(4)), k: 4}, true},
		{"tree", args{a: tree.Insert(int64(2)), k: 2}, true},
		{"tree", args{a: tree.Insert(int64(1)), k: 1}, true},
		{"tree", args{a: tree.Insert(int64(3)), k: 3}, true},
		{"tree", args{a: tree.Insert(int64(6)), k: 6}, true},
		{"tree", args{a: tree.Insert(int64(5)), k: 5}, true},
		{"tree", args{a: tree.Insert(int64(7)), k: 7}, true},
		{"tree", args{a: tree.Insert(int64(10)), k: 10}, true},
		{"tree", args{a: tree.Insert(int64(9)), k: 9}, true},
		{"tree", args{a: tree.Insert(int64(11)), k: 11}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.a.Search(int64(tt.args.k)) != tt.want {
				t.Errorf("BinaryTree: failed??? %v", tt.want)
			}
		})
	}

	t.Logf("%s", tree.String())

}
func TestSerializeDeserialize(t *testing.T) {

	tree := internal.NewBinaryTree()
	tree.Insert(8)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(10)
	tree.Insert(11)
	tree.Insert(9)

	trep := internal.Deserialize(tree.Serialize())

	if trep.Search(9) != true {
		t.Errorf("Failed")
	}

	if tree.Root.Left.Left.Data != trep.Root.Left.Left.Data {
		t.Error("Not valid\n")
		t.Fail()
	}
}
