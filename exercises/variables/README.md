# Variables

Variables in go are used to bind a value to a name, just like in other programming languages.

There are two main ways to define variables in go.

1. Using the `var` keyword at the package or function level
```aiignore
var x int
var x = 5
var x unint8 = 5
```

2. Using short assignment inside of a function
```aiignore
x := 5
```

## Constants

Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

Constants cannot be altered or re-assigned.

Constants should be named like variable in regard to capitalization (Capitalized const will be exported).

## Zero Value
Variables declared without an explicit initial value are given their zero value.

The zero value is:

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.