package main

import "fmt"

type Foo struct { 
  X int // weird to me to declare the name of the variable before the type/
  Y int
}


func main() {

  fmt.Println(Foo{1,2}) // print {1 2}

  // assign v to type Foo
  v := Foo{1, 2} // inialize X and Y in struct FOO
  // structs are accessed via a dot
  v.X = 4        // modify the "public" variable

  fmt.Println("v.X: ", v.X)

}
