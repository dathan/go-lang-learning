package challenges_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
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
	fmt.Printf("Len: %d Moving Right: %d\n", len, moveRight)
	// going from left to right N spaces
	for i := 1; i <= moveRight; i++ {
		removeThis = removeThis.Next
	}
	// deleteting the end (need to back up one)
	if removeThis != nil && removeThis.Next == nil && removeThis != head {
		removeThis = head
		moveRight := len - n - 1
		for i := 1; i <= moveRight; i++ {
			removeThis = removeThis.Next
		}
		removeThis.Next = nil
		return head
	}

	// 1->2->3->4->5: if n == 2; 5-2 == 3
	if removeThis != nil && removeThis.Next == nil && removeThis == head {
		return nil
	}

	//fmt.Printf("Removing node pos[%d] on len[%d]: %v\n", moveRight, len, removeThis)
	deleteNode(removeThis)
	return head

}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}

	n1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	n2 := &ListNode{1, nil}
	n3 := &ListNode{1, &ListNode{2, nil}}
	n4 := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"remove", args{head: n1, n: 2}, n1},
		{"remove", args{head: n2, n: 1}, nil},
		{"remove", args{head: n3, n: 1}, n3},
		{"remove", args{head: n4, n: 1}, n4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			} else {
				t.Log(spew.Sprintf("OK: %v\n", tt.want))
			}
		})
	}
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
