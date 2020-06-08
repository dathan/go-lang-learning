package challenges_test

import (
	"reflect"
	"testing"
)

/*
Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was `[1, 2, 3, 4, 5]`, the expected output would be `[120, 60, 40, 30, 24]`. If our input was `[3, 2, 1]`, the expected output would be `[2, 3, 6]`.

Follow-up: what if you can't use division?
*/
func TestListProduct(t *testing.T) {

	type args struct {
		a []int
		k []int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"productisok", args{[]int{1, 2, 3, 4, 5}, []int{120, 60, 40, 30, 24}}, true},
		{"productisok", args{[]int{3, 2, 1}, []int{2, 3, 6}}, true},
		{"productisbad", args{[]int{1, 4, 5, 7, 8, 10}, []int{}}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListProduct(tt.args.a, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

// two for loops to sum numbers is NxN == N^2
func ListProduct(a []int, answer []int) bool {

	var productArr []int = make([]int, len(a))

	for i, _ := range a {
		for j, _ := range a {
			if i == j { // do not product yourself.
				continue
			}

			if isset(productArr, i) == false {
				productArr[i] = a[j]
				continue
			}

			productArr[i] *= a[j]
		}
	}

	//spew.Dump(productArr, answer)

	for i, item := range answer {
		if item != productArr[i] {
			return false
		}
	}

	return true
}

func isset(arr []int, index int) bool {
	return (arr[index] > 1)
}
