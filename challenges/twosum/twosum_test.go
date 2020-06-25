package challenges_test

import (
	"reflect"
	"testing"
)

/*
 *
Given an array of integers, return indices of the two numbers such that they add up to a specific target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
*/
func twoSumBF(nums []int, target int) []int {

	ans := []int{}

	for i := 0; i < len(nums); i++ {
		//fmt.Printf("I: %d\n", i)
		for j := 0; j < len(nums); j++ {
			delta := nums[i] + nums[j]
			//fmt.Printf("\ti: %d j: %d nums[i]:%d nums[j]:%d sum: %d == target: %d\n", i, j, nums[i], nums[j], delta, target)
			if i == j {
				continue
			}

			if delta == target {

				ans = []int{i, j}
				return ans
			}
		}
	}

	return ans
}

func twoSumOpt(nums []int, target int) []int {

	ans := make([]int, 2)

	complimentMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {

		compliment := target - nums[i]
		if _, ok := complimentMap[compliment]; ok {
			ans[0] = complimentMap[compliment]
			ans[1] = i
			return ans
		}
		complimentMap[nums[i]] = i

	}

	return ans

}
func Test_twoSum(t *testing.T) {
	type args struct {
		input  []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"twosum", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
		{"twosum", args{[]int{3, 2, 4}, 6}, []int{1, 2}},
		{"twosum", args{[]int{3, 2, 3}, 6}, []int{0, 2}},
		{"twosum", args{[]int{0, 4, 3, 0}, 0}, []int{0, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumOpt(tt.args.input, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
