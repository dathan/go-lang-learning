package algo

import (
	"math/rand"
)

/**
Quicksort is a divide and conquer algorithm.
	Quicksort first divides a large array into two smaller sub-arrays:
		the low elements and the high elements.
	Quicksort can then recursively sort the sub-arrays.

The steps are:

Pick an element, called a pivot, from the array.
Partitioning: reorder the array so that all elements with values less than the
pivot come before the pivot, while all elements with values greater than the
pivot come after it (equal values can go either way). After this partitioning,
the pivot is in its final position. This is called the partition operation.
Recursively apply the above steps to the sub-array of elements with smaller
values and separately to the sub-array of elements with greater values.
The base case of the recursion is arrays of size zero or one, which are in order
by definition, so they never need to be sorted.

The pivot selection and partitioning steps can be done in several
different ways; the choice of specific implementation schemes greatly affects
the algorithm's performance.
*/
func QuickSort(a []int) []int {

	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// Pick a pivot
	pivotIndex := rand.Int() % len(a)

	// Move the pivot to the right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	// Pile elements smaller than the pivot on the left
	for i := range a {
		if a[i] < a[right] {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	// Place the pivot after the last smaller element
	a[left], a[right] = a[right], a[left]

	// Go down the rabbit hole
	QuickSort(a[:left])
	QuickSort(a[left+1:])

	return a
}
