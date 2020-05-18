package arrays

func Concat(a, b []int) []int {

	c := append(a, b...)

	return c

}
