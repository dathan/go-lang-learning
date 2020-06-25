package challenges_test

import (
	"fmt"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

/**
You are given an n x n 2D matrix representing an image.

Rotate the image by 90 degrees (clockwise).

Note:

You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Example 1:

Given input matrix =
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

rotate the input matrix in-place such that it becomes:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
Example 2:

Given input matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
],

rotate the input matrix in-place such that it becomes:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
*/

// need a swap
func swap(input []int, a, b int) {
	input[a], input[b] = input[b], input[a]
}

// when moving 90 degrees row 1 swaps with col N, row 2 swaps with col N-1,..
// so row 0: 1,2,3 : (0,0), (0,1), (0,2) maps to col 2: (0,2), (1,2), (2,2)
// In place rotate a N x N matrix by 90 degrees in a clockwise direction
func rotate(matrix [][]int) {
	//Consider the squares 1 by 1
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		// consider elements in groups of N in a current square
		for j := i; j < n-i-1; j++ {
			// store the current cell in temp variable
			temp := matrix[i][j]
			// move top left to top right
			temp, matrix[j][n-1-i] = matrix[j][n-1-i], temp         // 3:1
			temp, matrix[n-1-i][n-1-j] = matrix[n-1-i][n-1-j], temp // 9:3
			temp, matrix[n-1-j][i] = matrix[n-1-j][i], temp         // 7:9
			matrix[i][j] = temp                                     // 1:7
		}
	}
	fmt.Printf("Display\n")
	spew.Dump(matrix)
}

func Test_rotate(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
	}{
		{"basic", args{[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.matrix)
		})
	}
}
