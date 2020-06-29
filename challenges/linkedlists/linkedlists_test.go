package challenges_test

import (
	"reflect"
	"testing"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// moving the current node value to the next value and pointing next to next, the "GC" will clean up the object
func deleteNode(node *ListNode) {
	if node != nil && node.Next != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	len := 1
	d := head
	for d != nil && d.Next != nil {
		len++
		d = d.Next
	}

	if n > len {
		return head
	}

	removeThis := head
	moveRight := len - n
	// going from left to right N spaces
	for i := 1; i <= moveRight; i++ {
		removeThis = removeThis.Next
	}

	// 1->2->3->4->5: if n == 2; 5-2 == 3
	if removeThis != nil && removeThis.Next == nil {
		return nil
	}

	//fmt.Printf("Removing node pos[%d] on len[%d]: %v\n", moveRight, len, removeThis)
	deleteNode(removeThis)
	return head

}

func Test_deleteNode(t *testing.T) {
	type args struct {
		node *ListNode
	}

	n := &ListNode{4, &ListNode{5, &ListNode{1, &ListNode{9, nil}}}}

	tests := []struct {
		name string
		args args
	}{
		{"delete", args{n.Next}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.args.node)
		})
	}
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}

	n := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	n2 := &ListNode{1, nil}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"remove", args{head: n, n: 2}, n},
		{"remove", args{head: n2, n: 1}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
