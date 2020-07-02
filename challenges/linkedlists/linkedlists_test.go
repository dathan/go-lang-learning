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

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 1->2->2->1
func isPalindrome(head *ListNode) bool {

	fast := head
	slow := head
	stack := []int{}
	// via pointers cut the list in 1/2 put them in the stack in order
	for fast != nil && fast.Next != nil {
		fmt.Printf("FAST: %d SLOW: %d\n", fast.Val, slow.Val)
		stack = append(stack, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	fmt.Printf("\tSTACK: %v\n", stack)
	if fast != nil { // the list is cut in 1/2 it might be odd handle that casew
		slow = slow.Next
	}

	// the next to last is looking at the end
	for slow != nil {
		fmt.Printf("SLOW-c: %d == %d\n", slow.Val, stack[len(stack)-1])
		if slow.Val != stack[len(stack)-1] {
			return false
		}
		slow = slow.Next
		stack = stack[:len(stack)-1]
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	if len(res) == 1 {
		return true
	}
	if len(res) == 2 {
		return res[0] == res[1]
	}
	if len(res) == 3 {
		return res[0] == res[2]
	}
	if len(res)%2 == 0 {
		return helper(res[:len(res)/2], res[len(res)/2:])
	}
	return helper(res[:len(res)/2], res[len(res)/2+1:])
}

func helper(a, b []int) bool {
	for i, j := 0, len(a)-1; i < len(a); i, j = i+1, j-1 {
		if a[i] != b[j] {
			return false
		}
	}
	return true
}

/*
Merge two sorted linked lists and return it as a new sorted list.
The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	//fmt.Printf("L1: %s L2: %s\n", spew.Sdump(l1), spew.Sdump(l2))
	// use the dummy node method. We are moving each listNode to do the sort
	var dummyNode *ListNode = &ListNode{}

	// tail now points to the last result of dummy node
	tail := dummyNode
	for {

		// we iterated through the entire list node so that means the other list is the sort
		if l1 == nil {
			tail.Next = l2
			break
		}

		// same explination as above
		if l2 == nil {
			tail.Next = l1
			break
		}

		if l2.Val < l1.Val {
			tail.Next = l2
			l2 = l2.Next // move the pointer to the next element
		} else {
			tail.Next = l1
			l1 = l1.Next
		}

		//advance the tail
		tail = tail.Next
	}

	//fmt.Printf("tail: %s\n dummyNode: %s\n\tdummyNode.next: %s\n", spew.Sdump(tail), spew.Sdump(dummyNode), spew.Sdump(dummyNode.Next))
	return dummyNode.Next // remember dummyNode holds the 0 we prepended to the list
}

// 1->2->3->4->5->NULL want 5->4->3->2->1->NULL: 0: 2->3->4->5->1->NULL 1: 3->4->5->2->1
func reverseList(head *ListNode) *ListNode {

	// prev is the state of current after the swap
	var prev *ListNode = nil

	current := head      // current points to head
	for current != nil { // keep moving the list until nil is reached
		nextTemp := current.Next // 0: 2->3->4->5->nil, 1: 3->4->5->nil
		current.Next = prev      // 0: 1->nil 1: 2->1->nil
		prev = current           // 0: 1->nil 1: 2->1->nil
		current = nextTemp       // 0: 2->3->4->5->nil 1: 3->4->5->nil
	}

	return prev
}

func reverseListRecursive(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head // this makes current as head
	}
	var prev *ListNode = nil
	//fmt.Printf("HEAD: %s\n", spew.Sdump(head))
	prev = reverseListRecursive(head.Next) // stack keeps advancing the pointer then plays below
	// this reverses the list
	head.Next.Next = head
	head.Next = nil // this detaches the list element from next
	//fmt.Printf("LIST: %s\n\t HEAD PREV: %s", spew.Sdump(prev), spew.Sdump(head))
	return prev
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

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}

	n1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"reverseit", args{n1}, n1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			} else {
				t.Logf("reverseList() = %v", got)
			}
		})
	}
}

func Test_reverseListRecursive(t *testing.T) {
	type args struct {
		head *ListNode
	}

	n1 := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"reverseit", args{n1}, n1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseListRecursive(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseListRecursive() = %v, want %v", got, tt.want)
			} else {
				t.Logf("reverseListRecursive() = %v", got)
			}
		})
	}
}

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}

	l1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	l2 := &ListNode{1, &ListNode{3, &ListNode{5, nil}}}

	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"Want", args{l1, l2}, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}

	head := &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	headFail := &ListNode{1, &ListNode{2, nil}}
	headFail2 := &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{1, nil}}}}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1221", args{head}, true},
		{"12", args{headFail}, false},
		{"1121", args{headFail2}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
