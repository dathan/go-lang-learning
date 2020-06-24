package challenge_test

import (
	"reflect"
	"testing"
)

/*
Given an array, rotate the array to the right by k steps, where k is non-negative.

Follow up:

Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?


Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]


Constraints:

1 <= nums.length <= 2 * 10^4
It's guaranteed that nums[i] fits in a 32 bit-signed integer.
k >= 0
*/

func TestRotateArray(t *testing.T) {

	type args struct {
		input    []int
		rotateBy int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"rotateArray", args{input: []int{1, 2, 3, 4, 5, 6, 7}, rotateBy: 3}, []int{5, 6, 7, 1, 2, 3, 4}},
		{"rotateArray", args{input: []int{1, 2, 3, 4, 5, 6, 7}, rotateBy: 4}, []int{4, 5, 6, 7, 1, 2, 3}},
		{"rotateArray", args{input: []int{-1, -100, 3, 99}, rotateBy: 2}, []int{3, 99, -1, -100}},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if rotateArraySameSpace(tt.args.input, tt.args.rotateBy); !reflect.DeepEqual(tt.args.input, tt.want) {
				t.Errorf("rotateArray = %v, want %v\n", tt.args.input, tt.want)
			}
		})
	}
	/*
		input := []int{1, 2, 3, 4, 5, 6, 7}

		t.Logf("pre nums: %+v\n", input)

		rotateArray(&input, 3)

		t.Logf("post nums: %+v\n", input)
	*/
}

//passing a slice is passed by reference but append is making a copy thus operate on the value slice via a pointer of slices (which are just pointers "view" to an array) Space: O(N) + O(1)
func rotateArray(nums *[]int, k int) { //TODO benchmark and track allocations

	if k > len(*nums)-1 { // skip invalid loops
		return
	}

	i := len(*nums) - 1

	count := 0 // number of pops and unshits onto the array
	for count < k {
		// take off element at the end
		arraySet := (*nums)[i]
		*nums = append((*nums)[:i], (*nums)[i+1:]...) // pops the item off
		*nums = append([]int{arraySet}, (*nums)...)   // array_unshift (prepends to an array)
		count++
	}

}

func rotateArraySameSpace(nums []int, k int) { // bruteforce same space
	for i := 1; i <= k; i++ {
		previous := nums[len(nums)-1]
		for j := 1; j < len(nums); j++ {
			nums[j], previous = previous, nums[j]
		}
	}
}
