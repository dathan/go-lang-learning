package algo

/**
In computer science, binary search, also known as half-interval search,
logarithmic search, or binary chop, is a search algorithm that finds the
position of a target value within a sorted array.
Binary search compares the target value to the middle element of the array;
	if they are unequal,
		the half in which the target cannot lie is eliminated and
		the search continues on the remaining half until it is successful.
	If the search ends with the remaining half being empty,
		the target is not in the array.

Binary search runs in at worst logarithmic time, making O(log n) comparisons,
where n is the number of elements in the array,
the O is Big O notation, and log is the logarithm.
Binary search takes constant (O(1)) space, meaning that the space taken by
the algorithm is the same for any number of elements in the array.
Although specialized data structures designed for fast searching—such as
hash tables—can be searched more efficiently, binary search applies to a
wider range of problems.
*/
func BinarySearch(target_map []int, value int) int {

	start_index := 0
	end_index := len(target_map) - 1

	for start_index <= end_index {

		median := (start_index + end_index) / 2

		if target_map[median] < value {
			start_index = median + 1
		} else {
			end_index = median - 1
		}

	}

	if start_index == len(target_map) || target_map[start_index] != value {
		return -1
	} else {
		return start_index
	}

}
