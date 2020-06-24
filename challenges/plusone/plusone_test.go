package challenges_test

import (
	"reflect"
	"testing"
)

/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list,
and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.
Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.

*/
func plusone(input []int) []int {

	var ret []int = make([]int, 0)

	i := len(input) - 1
	if i < 0 {
		return []int{1} // overflow / carryover
	}
	//fmt.Printf("INPUT: %v I: %d\n", input, i)
	p := input[i] + 1
	if p >= 10 {
		p = p - 10
		return append(plusone(input[:i]), p)
	}

	input[i] = input[i] + 1
	ret = input
	return ret

}

func Test_plusone(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"basic", args{[]int{1, 2, 3}}, []int{1, 2, 4}},
		{"overflow", args{[]int{9, 9, 9}}, []int{1, 0, 0, 0}},
		{"bounds", args{[]int{9, 8, 9}}, []int{9, 9, 0}},
		{"basic", args{[]int{4, 3, 2, 1}}, []int{4, 3, 2, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusone(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("plusone() = %v, want %v", got, tt.want)
			}
		})
	}
}
