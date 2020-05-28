package arrays

import "fmt"

/**
 * Slices are just views of arrays
 */

// learn more about the internals @ http://golang.org/doc/articles/slices_usage_and_internals.html
func SliceExample() {
	p := []int{2, 3, 5, 7, 11, 13}

	fmt.Println("p is a slice of an int array with 6 elements == ", p)
	fmt.Println("len == ", len(p))

	fmt.Println("now we can \"re-slice\", creating a new slice value pointer to the same array")
	fmt.Println("p[1:4] ==", p[1:4])

	fmt.Println("Now we will show a type of ALLOC called make")
	fmt.Println("The \"make\" function works by allocating a zero array and returning a slice that refers to that array")
	a := make([]int, 5) //len(a) == 5
	fmt.Println("len(a) == ", len(a))

	//
	// iterating an array
	//

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)

	}

	fmt.Println("We used the range operator to go through the array and return a key => value ordered by the array")

	/**
		 * note the equivlent of iterating through a map and just look at the values
	   * e.g. foreach($data as $row) { }
	   * is for _, value := range data { }
	*/

}
