Pointers are used for Go.

*T is a pointer to a T value

The & operator generates a pointer to its operand

i := 42
p = &i  p now points to i

*p = 21 // set i to 21 through the pointer p


`Go does not have pointer arithmetic`

### Dereferencing to get to the value example with an Interface type.

token *pkg.Token // token is pointing to a pkg.Token object

```
(*token.PublicType.(*pkg.PublicTypeInterfaceImpl))
type PublicType interface { MethodToImpl() error }

type Token {

    PublicTypeVar PublicType    

}


```

`*token.PublicTypeVar` // dereference token.PublicTypeVar then cast the PublicType interface into its typed value by casting to a derefernced value of (*pkg.PublicTypeInterfaceImpl) since PublicTypeVar is an interface





