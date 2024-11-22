# Arrays

The type `[n]T` is an array of n values of type T.

The expression
```
var a [10]int
```
declares a variable `a` as an array of ten integers.

Arrays can be initialized in a few different ways:
```aiignore
var a [3]int
var a [3]int{0,1,2}
var a [...]int{0,1,2}
```

An array's length is part of its type (fixed at compile time), so arrays cannot be resized.

### Comparable

Arrays are comparable.
```aiignore
a := [3]int{0, 1, 2}
b := [3]int{0, 1, 2}
fmt.Println(a == b)  // true
```
### Passed by

Arrays are passed by value so all elements are copied!