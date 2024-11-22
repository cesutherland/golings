# Functions

A function can take zero or more arguments.

The type comes after the variable name for parameters.

When two or more consecutive named function parameters share a type, you can omit the type from all but the last.
```aiignore
func add(x int, y int)

or 

func add(x, y int)
```

A function can return any number of results.
```aiignore
func swap(x, y string) (string, string) {
	return y, x
}
```

### Named Return Values
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. **They can harm readability in longer functions.**

```aiignore
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
```