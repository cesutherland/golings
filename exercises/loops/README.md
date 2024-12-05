# Loops

## For

Go has only one looping construct, the `for` loop.

The basic `for` loop has three components separated by semicolons:

* the init statement: executed before the first iteration

* the condition expression: evaluated before every iteration

* the post statement: executed at the end of every iteration

The init statement will often be a short variable declaration, and the variables declared there are visible only in 
the scope of the `for` statement.

The loop will stop iterating once the boolean condition evaluates to `false`.

```aiignore
func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding the three components 
of the `for` statement and the braces { } are always required.

### While Loops

**The init and post statements are optional.** As such, a while loop can be written in the following manner.

```aiignore
func main() {
	sum := 1
	for ; sum < 1000; {  // OR just `for sum < 1000 {`
		sum += sum
	}
	fmt.Println(sum)
}
```

## Range

The range form of the for loop iterates over a slice, array, or map.

When ranging over an array or slice, two values are returned for each iteration. 
The first is the index, and the second is a copy of the element at that index.

When ranging over a map, two values are returned for each iteration.
The first is the key, and the second is a copy of the value.

You can skip the index, key, or value by assigning to _.

```aiignore
list := make([]int, 10)
for i, _ := range list {
    // do something with the index i
}

for _, value := range list {
    // do something with a copy of the value
}
```

If you only want the index or key, you can omit the second variable.

`for i := range list`