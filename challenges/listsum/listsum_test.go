package challenges_test

import (
	"reflect"
	"testing"
)

/*
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.

Bonus: Can you do this in one pass?
*/
func TestListSum(t *testing.T) {

	type args struct {
		a []int
		k int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"sumisok", args{[]int{1, 4, 5, 7, 8, 10}, 17}, true},
		{"sumisok", args{[]int{1, 4, 5, 7, 8, 10}, 1270}, false},
		{"sumisok", args{[]int{1, 4, 5, 7, 8, 10}, 18}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListSum(tt.args.a, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersection() = %v, want %v", got, tt.want)
			}
		})
	}
}

// two for loops to sum numbers is NxN == N^2
func ListSum(a []int, k int) bool {

	for i, item := range a {
		for j, itemagain := range a {
			if i == j { // do not sum yourself.
				continue
			}
			if (item + itemagain) == k {
				return true
			}
		}
	}

	return false
}
