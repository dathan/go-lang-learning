package challenges_test

import (
	"reflect"
	"testing"
)

/*
Given two arrays, write a function to compute their intersection.

Example 1:

Input: nums1 = [1,2,2,1], nums2 = [2,2]
Output: [2,2]
Example 2:

Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
Output: [4,9]
Note:

Each element in the result should appear as many times as it shows in both arrays.
The result can be in any order.
Follow up:

What if the given array is already sorted? How would you optimize your algorithm?
What if nums1's size is small compared to nums2's size? Which algorithm is better?
What if elements of nums2 are stored on disk, and the memory is limited such that you cannot load all elements into the memory at once?
*/

func intersectBF(nums1 []int, nums2 []int) []int {

	var ans []int = []int{}

	for _, elm := range nums1 {

		for i := 0; i < len(nums2); i++ {
			if elm == nums2[i] {
				ans = append(ans, elm)
				nums2 = append(nums2[:i], nums2[i+1:]...)
			}
		}
	}
	return ans

}

func intersect(nums1 []int, nums2 []int) []int {
	var intersectMap map[int]int = make(map[int]int)
	for _, elm := range nums1 {
		intersectMap[elm]++
	}

	var ans []int
	for _, elm := range nums2 {
		if e, ok := intersectMap[elm]; ok && e > 0 {
			ans = append(ans, elm)
			intersectMap[elm]--
		}
	}
	return ans
}

func Test_intersect(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"ok", args{[]int{1, 2, 2, 1}, []int{2, 2}}, []int{2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
