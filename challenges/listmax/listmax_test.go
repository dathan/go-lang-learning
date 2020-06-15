package challenges_test

import (
	"reflect"
	"testing"
)

/*
Given an array of integers and a number k, where 1 <= k <= length of the array,
compute the maximum values of each subarray of length k.

For example, given array = [10, 5, 2, 7, 8, 7] and k = 3, we should get: [10, 7, 8, 8], since:

10 = max(10, 5, 2)
7 = max(5, 2, 7)
8 = max(2, 7, 8)
8 = max(7, 8, 7)
Do this in O(n) time and O(k) space. You can modify the input array in-place and you do not need to store the results. You can simply print them out as you compute them.
*/
// return the max of the sub arrays note k < size of the array
func listmax(input []int, k int) []int {
	//fmt.Printf("LEN: %d K: %d\n", len(input), k)
	iterations := len(input) + 1 - k
	subArray := make([]int, iterations)
	for i := 0; i < iterations; i++ {
		maxElm := input[i]
		iteration := 0
		for j := i; j < (i + k); j++ { // this is constant since its less then the array
			iteration++
			searchElm := j
			if j >= len(input) { // since K can == the len(input)
				panic("Should not go past the array")
			}
			if input[searchElm] > maxElm {
				maxElm = input[searchElm]
			}
			//fmt.Printf("Rotation: %d Looking at i: %d j: %d input[i]: %d input[j]: %d MAX: %d I: %d \n", iteration, i, j, input[i], input[searchElm], maxElm, iterations)
		}
		subArray[i] = maxElm
	}
	//fmt.Printf("MAX: %v\n", subArray)
	return subArray

}

func max(a []int) int {

	max := 0
	for elm := range a {
		if max < elm {
			max = elm
		}
	}
	return max
}

// what helped me was drawing this out on paper to see the number of iterations you can do within an array where k <= length
func Test_listmax(t *testing.T) {
	type args struct {
		input []int
		k     int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"listmax", args{input: []int{10, 5, 2, 7, 8, 7}, k: 3}, []int{10, 7, 8, 8}},
		{"listmax", args{input: []int{10, 5, 2, 7, 8, 7}, k: 5}, []int{10, 8}},
		{"listmax", args{input: []int{10, 5, 2, 7, 8, 7}, k: 6}, []int{10}},
		{"listmax", args{input: []int{10, 5, 2, 7, 8, 7}, k: 2}, []int{10, 5, 7, 8, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := listmax(tt.args.input, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("listmax() = %v, want %v", got, tt.want)
			}
		})
	}
}
