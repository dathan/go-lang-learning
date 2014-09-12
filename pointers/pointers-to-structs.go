package main

import "fmt"

type Foo struct {
	X int
	Y int
}

func main() {

	v := Foo{1, 2} // init the Foo Struct
	p := &v        // p now points to v
	p.X = 1e9      // shorthand to assign 1 x 10^9
	fmt.Println("p changed v to : ", v)

}
