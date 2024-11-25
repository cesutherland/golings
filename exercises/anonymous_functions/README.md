## Anonymous Functions

Go supports anonymous functions, which can form closures.
Anonymous functions are useful when you want to define a function inline without having to name it.

### Basic Anonymous Function

An anonymous function can be created as such:
```aiignore
sayHello := func() {
    fmt.Printf("Hello %s\n", n)
}

sayHello("World!") // Hello World!
```

### Closure

A closure is when a function inside another function closes over one or more 
local variables of the outer function.

**The inner function gets a reference to the outer function's variables**

```aiignore
func intSeq() func() int {  // returns a function that returns an int
    i := 0  // i is closed over in the returned func
    return func() int {
        i++
        return i
    }
}

// this can be called like such
nextInt := intSeq()

fmt.Println(nextInt())  // 1
fmt.Println(nextInt())  // 2
fmt.Println(nextInt())  // 3
```