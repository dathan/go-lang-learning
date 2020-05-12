package arrays

// given two lists of int arrays; a,b find the intersection of a and b
func Intersection(a, b []int) []int {

	intersection := make(map[int]bool)

	for _, item := range a {
		intersection[item] = true
	}

	var c []int
	for _, item := range b {

		if _, ok := intersection[item]; ok {
			c = append(c, item)
		}

	}

	return c

}
