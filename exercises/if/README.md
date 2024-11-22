# If

The `if` statement is used for flow control like other languages.

It can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the if.
```aiignore
func pow(x, n, lim float64) float64 {
	// v can be used within the if block
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}
```