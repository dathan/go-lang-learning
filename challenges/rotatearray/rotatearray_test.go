package challenge_test

import (
	"fmt"
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

	input := []int{1, 2, 3, 4, 5, 6, 7}

	t.Logf("pre nums: %+v\n", input)

	rotateArray(input, 3)

	t.Logf("post nums: %+v\n", input)

}

//passing a slice
func rotateArray(nums []int, k int) {
	if k > len(nums)-1 {
		return
	}
	// traverse the array backwards
	i := len(nums) - 1
	count := 0
	for count < k {
		// take off element at the end
		arraySet := nums[i]
		//fmt.Println(arraySet)
		nums = append(nums[:i], nums[i+1:]...) // pops the item off
		//fmt.Println(nums)
		nums = append([]int{arraySet}, nums...) // array_unshift (prepends to an array)
		//fmt.Println(nums)
		count++

	}
	fmt.Println(nums)

}
