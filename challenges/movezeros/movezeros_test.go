package challenges_test

import (
	"reflect"
	"testing"
)

/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative
order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/
func moveZeroes(nums []int) {
	// there is no reason to start from 0 each time for the inner loop,
	// do single swap, keep a pointer to each elment getting moved to the right of its position if 0
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for ; j < len(nums); j++ {
				if nums[j] != 0 && j > i {
					nums[i], nums[j] = nums[j], nums[i] // move zeros to the end by reference without copies
					break
				}
			}
		}
	}
}

//Make an array of Types and "Objectify" the solution above
type MoveZeroes []int

//MoveZeroesTo the End
func (nums MoveZeroes) Rearrange() {
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for ; j < len(nums); j++ {
				if nums[j] != 0 && j > i {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}

func TestMoveZeroes_Rearrange(t *testing.T) {
	tests := []struct {
		name string
		nums MoveZeroes
		want MoveZeroes
	}{
		{"ok", MoveZeroes{0, 1, 0, 3, 12}, MoveZeroes{1, 3, 12, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.nums.Rearrange()
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("findSignatures() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
