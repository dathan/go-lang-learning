package backtracing

import "testing"

/*
The following solution does not use memoization. It is a pure brute force solution.
nums : [1,5,5,11]
Total sum is 22 and we need to find if is there any combination possible that sum to 11(22/2).

How do we brute force it?
Ans: Generate all the possible combinations

**A suggestion: Solve problems tagged with Backtracking to understand this kind of recursion. **
*/

func canPartition(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	total := 0
	for _, val := range nums {
		total += val
	}

	if total%2 != 0 {
		return false
	}

	target := total / 2
	return backTrack(nums, target, 0)
}

func backTrack(nums []int, target int, idx int) bool {
	if idx == len(nums) || target < 0 {
		return false
	}
	if target == 0 {
		return true
	}
	return backTrack(nums, target-nums[idx], idx+1) || backTrack(nums, target, idx+1)
}

func Test_canPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"OK", args{[]int{1, 5, 5, 11}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); got != tt.want {
				t.Errorf("canPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
