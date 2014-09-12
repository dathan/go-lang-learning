Go does not have classes, although you can simulate the look and feel by
defining methods on struct types


Some Notes:

Interfaces are cool and confusing. You don't have to declare a "method" as
implements the Interface.

For instance fmt package interface Stringer just including that package and
implementing a function on your struct Person

type Person struct {
  Name String
  Age int
}

func (p Person) String() string { // implement fmt.Stringer's String Interface since it defines the interface
  return fmt.Sprintf("%v .....", p.Name, p.Age)
} 

func main() {

  fmt.Println(a,z); // will use String to print the value of a, z
}
