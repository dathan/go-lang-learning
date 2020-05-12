package arrays

import "fmt"

type Vertex struct {
	Lat, Lon float64
}

var m map[string]Vertex

func MapExample() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{40.64, -74.39}
	fmt.Println(m["Bell Labs"])

	// if "DOES NOT EXIST" is in the map ok is true and v has that value
	v, ok := m["DOES NOT EXIST"] // really cool syntax kind of like elvis syntax
	fmt.Println("The value v : ", v, "Present ?", ok)
}
