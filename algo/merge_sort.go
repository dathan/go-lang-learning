package algo

/*
Conceptually, a merge sort works as follows:

Divide the unsorted list into n sublists, each containing 1 element
(a list of 1 element is considered sorted).
Repeatedly merge sublists to produce new sorted sublists until there is only 1
sublist remaining. This will be the sorted list.
*/

func merge(l, r []int) []int {
	// make an array to returned the merged array
	ret := make([]int, 0, len(l)+len(r))
	// while left and right is > 0
	for len(l) > 0 || len(r) > 0 {

		if len(l) == 0 {
			return append(ret, r...) // its just the right
		}

		if len(r) == 0 {
			return append(ret, l...) // its just the left
		}

		if l[0] <= r[0] {
			ret = append(ret, l[0]) // move the single element to ret
			l = l[1:]               // l now pops off 0 element
		} else {
			ret = append(ret, r[0])
			r = r[1:] // right now pops off the 0 element
		}
	}
	return ret
}

/**
 * MergeSort sort left and right then merge
 */
func MergeSort(s []int) []int {
	if len(s) <= 1 { // single elemented is considered sorted
		return s
	}
	n := len(s) / 2       // find the mid point
	l := MergeSort(s[:n]) // sort the left
	r := MergeSort(s[n:]) // sort the right
	return merge(l, r)    // merge left and right
}
