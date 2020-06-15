package challenge_test

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

/**

Given a sorted list
[-5, -3, -2, 1, 2, 3, 4]

in O(n) time, square the array in the same space and produce a sorted array;

WHat I am doing here is going through the array backwards, testing to see if negative numbers are greater then the right number.
If so increment the left to move the "pointer" and store the result backwards, otherwise keep moving right squaring.
Think of the the array being moved from both the left and the right as the single traversal happens.
*/

func TestSortIntSquare(t *testing.T) {

	in := []int{-5, -3, -2, 1, 2, 3, 4}
	out := make([]int, len(in))
	inputLen := len(in) - 1
	left := 0
	right := inputLen
	for i := inputLen; i >= 0; i-- {

		fmt.Printf("left:[%d] %d right:[%d] %d i: [%d] %d\n", left, in[left], right, in[right], i, in[i])
		if abs(in[left]) > in[right] {
			out[i] = abs(in[left]) * abs(in[left])
			left++
		} else {
			out[i] = in[right] * in[right]
			right--
		}
	}
	spew.Dump(in, out)

}

func abs(i int) int {
	if i < 0 {
		i = i * -1
	}
	return i
}
