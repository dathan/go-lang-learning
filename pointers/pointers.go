package main

import "fmt"

func main() {
  
  i, j := 6, 9


  p := &i // point to i -- "indirecting"
  fmt.Println(*p) // read i through the pointer p (display the value of i) - "dereferencing"
  fmt.Printf("i : %v, j: %v : ( *p + i = %v )", i,j, (*p+j))
}
