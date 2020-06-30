package challenges_test

import (
	"fmt"
	"reflect"
	"testing"
)

/*
Contiguous Subarrays
You are given an array a of N integers. For each index i, you are required to determine the number of contiguous subarrays that fulfills the following conditions:
The value at index i must be the maximum element in the contiguous subarrays, and
These contiguous subarrays must either start from or end on index i.
Signature
int[] countSubarrays(int[] arr)
Input
Array a is a non-empty list of unique integers that range between 1 to 1,000,000,000
Size N is between 1 and 1,000,000
Output
An array where each index i contains an integer denoting the maximum number of contiguous subarrays of a[i]
Example:
a = [3, 4, 1, 6, 2]
output = [1, 3, 1, 5, 1]
Explanation:
For index 0 - [3] is the only contiguous subarray that starts (or ends) with 3, and the maximum value in this subarray is 3.
For index 1 - [4], [3, 4], [4, 1]
For index 2 - [1]
For index 3 - [6], [6, 2], [1, 6], [4, 1, 6], [3, 4, 1, 6]
For index 4 - [2]
So, the answer for the above input is [1, 3, 1, 5, 1]
*/
func countSubarrays(arr []int) []int {
	result := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		numberOfArrays := 1

		//Move from Right to Left
		for left := i - 1; left >= 0; left-- {
			if arr[left] < arr[i] {
				numberOfArrays++
				continue
			}
			break
		}

		//Move from Left to Right
		for right := i + 1; right < len(arr); right++ {
			if arr[right] < arr[i] {
				numberOfArrays++
				continue
			}

			break
		}

		result[i] = numberOfArrays
	}

	return result
}

func printSubArrays(arr []int, start, end int) []int {

	// stop if reach the end of the array
	if end == len(arr) {
		return nil
	}

	// increment the end point and start from 0
	if start > end {
		return printSubArrays(arr, 0, end+1)
	}

	fmt.Print("[")
	for i := start; i < end+1; i++ {
		fmt.Printf("%d", arr[i])
		if i != end {
			fmt.Printf(",")
		}
	}
	fmt.Printf("]\n")
	return printSubArrays(arr, start+1, end)

}

func Test_countSubarrays(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

		{"incrementtoanswer", args{[]int{3, 4, 1, 6, 2}}, []int{1, 3, 1, 5, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubarrays(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printSubArrays(t *testing.T) {
	type args struct {
		arr   []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"TestSubArray", args{[]int{1, 2, 3}, 0, 0}, nil},
		{"TestSubArray", args{[]int{3, 4, 1, 6, 2}, 1, 0}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := printSubArrays(tt.args.arr, tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printSubArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
