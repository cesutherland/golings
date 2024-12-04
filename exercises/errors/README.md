## Errors

Go programs express error state with error values.

The error type is a built-in **interface** similar to `fmt.Stringer`:

```
type error interface {
    Error() string
}
```

(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an `error` value, and calling code should handle errors by testing whether the error equals `nil`.

```
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}

fmt.Println("Converted integer:", i)
```

A `nil` `error` denotes success; a non-nil `error` denotes failure.

### Constructing Errors

Errors can be constructed on the fly using Go’s built-in `errors` or `fmt` packages.

The following function uses the `errors` package to return a new `error` with a static error message
```aiignore
import "errors"

func DoWork() error {
    return errors.New("something broke")
}
```

The `fmt` package can be used to add dynamic data to the error, such as an `int`, `string`, or another `error`
```aiignore
import "fmt"

func DoWork() error {
    return fmt.Errorf("something broke: %s", "something")
}
```

### Sentinel Errors

We can use sentinel errors to explicitly check for an expected error.

```aiignore
var ErrMyError = errors.New("something broke")

func DoWork() error {
    return ErrMyError
}

func main() {
    err := DoWork()
    if err != nil {
        switch {
        case errors.Is(err, ErrMyError):
            // do something with the ErrMyError
        }
    ...
    }
}
```

### Custom Errors

Sometimes it is helpful for errors to carry additional data fields, or maybe the error’s message should populate itself with dynamic values when it’s printed.

You can do that in Go by implementing custom errors type.

```aiignore
type MyError struct {
    Var1 string
    Msg  string
}

func (m MyError) Error() {
    fmt.Sprintf("there was an error with %s and %s: (%s)", e.Var1, e.Var2, e.Msg)
}

func Do() error {
    return &MyError{"var1", "var2", "ahhhh"}
}
```

### Panic

The usual way to report an error to a caller is to return an error as an extra return value as above. 
But what if the error is unrecoverable? Sometimes the program simply cannot continue.

For this purpose, there is a built-in function `panic` that in effect creates a run-time error that will stop the program.
The function takes a single argument of arbitrary type—often a string—to be printed as the program dies. It's also a way to indicate that something impossible has happened, such as exiting an infinite loop.

```aiignore
var user = os.Getenv("USER")

func init() {
    if user == "" {
        panic("no value for $USER")
    }
}
```

### Recover

When `panic` is called, including implicitly for run-time errors such as indexing a slice out of bounds or failing a type assertion, it immediately stops execution of the current function and begins unwinding the stack of the goroutine, running any deferred functions along the way. 
If that unwinding reaches the top of the goroutine's stack, the program dies. However, it is possible to use the built-in function `recover` to regain control of the goroutine and resume normal execution.

A call to `recover` stops the unwinding and returns the argument passed to `panic`. Because the only code that runs while unwinding is inside deferred functions, `recover` is only useful inside deferred functions.

One application of `recover` is to shut down a failing goroutine inside a server without killing the other executing goroutines.

```aiignore
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)
    }
}

func safelyDo(work *Work) {
    defer func() {
        if err := recover(); err != nil {
            log.Println("work failed:", err)
        }
    }()
    do(work)
}
```

In this example, if `do(work)` panics, the result will be logged and the goroutine will exit cleanly without disturbing the others. There's no need to do anything else in the deferred closure; calling recover handles the condition completely.