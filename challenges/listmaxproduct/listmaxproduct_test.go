package listmaxproduct

import (
	"sort"
	"testing"
)

/**

Given a list of integers, return the largest product that can be made by multiplying any three integers.

For example, if the list is [-10, -10, 5, 2], we should return 500, since that's -10 * -10 * 5.

You can assume the list has at least three integers.
*/

type sortedInput []int

func (a sortedInput) Len() int {
	return len(a)
}

func (a sortedInput) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a sortedInput) Less(i, j int) bool {

	return a[i] < a[j]
}

func maxproduct(input []int) int {

	if len(input) < 3 {
		return 0
	}

	sort.Sort(sortedInput(input))

	maxProduct1 := input[len(input)-1] * input[len(input)-2] * input[len(input)-3]
	maxProduct2 := input[0] * input[1] * input[len(input)-1]

	if maxProduct1 > maxProduct2 {
		return maxProduct1
	}

	return maxProduct2
}

func Test_maxproduct(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"OK", args{[]int{-10, -10, 5, 2}}, 500},
		{"OK", args{[]int{-5, -5, -4, 2}}, 50},
		{"OK", args{[]int{-10, 8, 9, 10}}, 720},
		{"OK", args{[]int{2, 2, 2, 5}}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxproduct(tt.args.input); got != tt.want {
				t.Errorf("maxproduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
