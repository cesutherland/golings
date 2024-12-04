# Slices

Arrays have their place, but they’re a bit inflexible, so you don’t see them too often in Go code. Slices, though, are everywhere. They build on arrays to provide great power and convenience.

The type specification for a slice is `[]T`, where `T` is the type of the elements of the slice. Unlike an array type, a slice type has no specified length.

A slice literal is declared just like an array literal, except you leave out the element count:

`letters := []string{"a", "b", "c", "d"}`

A slice can be created with the built-in function called `make`, which has the signature,

`func make([]T, len, cap) []T`

where `T` stands for the element type of the slice to be created. The make function takes a type, a length, and an optional capacity. When called, make allocates an array and returns a slice that refers to that array.

```
var s []byte
s = make([]byte, 5, 5)
// s == []byte{0, 0, 0, 0, 0}
```
When the capacity argument is omitted, it defaults to the specified length. Here’s a more succinct version of the same code:

```
s := make([]byte, 5)
```

The length and capacity of a slice can be inspected using the built-in len and cap functions.

```
len(s) == 5
cap(s) == 5
```

See the [Slices package](https://pkg.go.dev/slices) for further documentation.

### Zero Value
The zero value of a slice is `nil`. The len and cap functions will both return 0 for a nil slice.

### Slicing Arrays

A slice can also be formed by “slicing” an existing slice or array. Slicing is done by specifying a half-open range with two indices separated by a colon. For example, the expression `b[1:4]` creates a slice including elements 1 through 3 of `b` (the indices of the resulting slice will be 0 through 2).

```
b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b
```
The start and end indices of a slice expression are optional; they default to zero and the slice’s length respectively:

```
// b[:2] == []byte{'g', 'o'}
// b[2:] == []byte{'l', 'a', 'n', 'g'}
// b[:] == b
```

This is also the syntax to create a slice given an array:
```
x := [3]int{0, 1, 2}
s := x[:] // a slice referencing the storage of x
```

### Appending to a Slice

It is common to append new elements to a slice, and so Go provides a built-in append function. The documentation of the built-in package describes `append`.

`func append(s []T, vs ...T) []T`

The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

**If the backing array of `s` is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.**

### Comparable

Slices are NOT comparable
```aiignore
a := []int{0, 1, 2}
b := []int{0, 1, 2}
fmt.Println(a == b) // invalid operation: a == b (slice can only be compared to nil)
```

### Passed by

Slices can be treated as if passed by reference, though be cautious if appending in the target function/method.