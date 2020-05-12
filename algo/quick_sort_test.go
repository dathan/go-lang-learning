package algo

import (
	"fmt"
	"testing"
)

func TestQuickSort(t *testing.T) {

	to_sort := []int{90383, 444, 3, 1, 0, 17483, 3232, 996, -3, -4}

	sorted_list := QuickSort(to_sort)

	fmt.Printf("%v\n", sorted_list)

}
