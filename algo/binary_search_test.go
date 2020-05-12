package algo

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	values := []int{1, 2, 3, 4, 5, 6, 7}
	index := BinarySearch(values, 3)

	if index == -1 {
		fmt.Printf("Bsearch missing %+v\n", index)
	} else {
		fmt.Printf("BinarySearch: Index: %d => Found %d\n", index, values[index])
	}

}
