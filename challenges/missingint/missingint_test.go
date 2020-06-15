package challenge_test

import (
	"fmt"
	"math/rand"
	"testing"
)

/*
Given an array of integers, find the first missing positive integer in linear time and constant space.
In other words, find the lowest positive integer that does not exist in the array.
The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place.
*/

func TestMissingInt(t *testing.T) {

	input := []int{3, 4, -1, 1}
	input = quicksort(input) // this makes it O(nlogn)

	// there are 3 O(n) sort algorithms but they must meet conditions
	// countingsort (works only on positive ints) https://play.golang.org/p/ttRVpPXW9A
	// radix sort (https://play.golang.org/p/Xmy0RPVXAv)
	// (Counting sort and radix sort assume that the input consists of integers in a small range)
	// bucket sort (+ insertion sort O(n)) https://play.golang.org/p/oTqMLk_lKZ
	// http://www.personal.kent.edu/~rmuhamma/Algorithms/MyAlgorithms/Sorting/linearTimeIntro.htm#:~:text=There%20are%20sorting%20algorithms%20that,radix%20sort%20and%20bucket%20sort.

	sizeofInput := len(input)
	want := 0
	for i, in := range input {
		if in < 0 {
			continue
		}

		nextDigit := in + 1
		nextPos := i + 1
		if nextPos > sizeofInput {
			nextPos = sizeofInput
		}

		if input[nextPos] != nextDigit {
			want = nextDigit
			break
		}
	}

	fmt.Printf("Want: %d\n", want)

}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
