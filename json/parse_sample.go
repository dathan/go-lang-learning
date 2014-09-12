package main

import "encoding/json"
import "fmt"

type Crates struct {
	Bundle int      `json:"bundle"`
	Fruits []string `json:"fruits"`
}

//https://gobyexample.com/json
func main() {
	//
	// Marshal encodes
	//
	bolB, _ := json.Marshal(true) // returns ([] byte, error)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))
	// slices
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	//
	// Unmarshal decodes
	//

	json_str := `{"num":6.13,"strs":["a","b"]}`
	byt := []byte(json_str)

	// TODO find out what 'interface{}' in this syntactic sugar means
	var dat map[string]interface{} // variable for the JSON package to put decoded data - it will hold a map of strings to arbitrary data types

	// now do the decode
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64) // cast the string to a float
	fmt.Println(num)

	// to access nested arrays we will need nested casts
	strs := dat["strs"].([]interface{}) // weird why are we casting from this interface
	str1 := strs[0].(string)
	fmt.Println(str1)

	// But what is really cool is this -> JSON directly to objects
	json_str2 := `{"bundle" : 1, "fruits": ["apple", "peach"]}`
	res := &Crates{}
	json.Unmarshal([]byte(json_str2), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])
}
