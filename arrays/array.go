package main

import "fmt"

func main() {
	var a [2]string // two values of Strings
	a[0] = "Dathan"
	a[1] = "CTO"

	fmt.Println(a[0], a[1]) // prints as two seprate words
	fmt.Println(a)          // prints as array format

	fmt.Println("Slices are next - they point to an array of values and also includes length")
}
