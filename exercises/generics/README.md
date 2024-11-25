## Generics

### Generic Type Parameters

Go functions can be written to work on multiple types using type parameters. The type parameters of a function appear between brackets, before the function's arguments.

`func Index[T comparable](s []T, x T) int`

This declaration means that `s` is a slice of any type `T` that fulfills the built-in constraint comparable. `x` is also a value of the same type.

### Generic Types

In addition to generic functions, Go also supports generic types. A type can be parameterized with a type parameter, which could be useful for implementing generic data structures.

```aiignore
// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}
```

### Constraints

Each type parameter has a type constraint that acts as a kind of meta-type for the type parameter. Each type constraint specifies the permissible type arguments that calling code can use for the respective type parameter.

While a type parameter’s constraint typically represents a set of types, at compile time the type parameter stands for a single type – the type provided as a type argument by the calling code. 
If the type argument’s type isn’t allowed by the type parameter’s constraint, the code won’t compile.

Constraints can come from built-in types, a union of types, an interface of unions, or from the 
[constraints package](golang.org/x/exp/constraints)

```aiignore
func PrintMe[T int | string](thing T) {
	fmt.Println(thing)
}

func main() {

	x := 10
	y := "Hello"
	z := 3.14

	PrintMe(x) // OK
	PrintMe(y) // OK
	PrintMe(z) // float64 does not satisfy int | string (float64 missing in int | string)
}
```

```aiignore
type Number interface {
	int | float64
}

func PrintMe[T Number](thing T) {
	fmt.Println(thing)
}

func main() {

	x := 10
	y := "Hello"
	z := 3.14

	PrintMe(x) // OK
	PrintMe(y) // string does not satisfy Number (string missing in int | float64)
	PrintMe(z) // OK
}
```