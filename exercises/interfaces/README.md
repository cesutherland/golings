## Interfaces

An interface type is defined as a set of method signatures.

A value of interface type can hold any value that implements those methods.

In Go, most interfaces have an `er` suffix: `Reader`,`Writer`, `Closer`, etc. That’s because, in Go, interfaces expose behavior, and their names point to that behavior.

```aiignore
type Stringer interface {  // Stringer interface from fmt package
    String() string
}
```

A type implements an interface by implementing its methods. There is no explicit declaration of intent, no "implements" keyword.

**All methods must be present to satisfy the interface and all methods for a given type must be declared in the same package where the type is declared**

Implicit interfaces decouple the definition of an interface from its implementation, which could then appear in any package without prearrangement.

```aiignore
type MyType struct{}

func (MyType) String() string {  // implicit implement of Stringer interface
	return "I'm a MyType"
}

func main() {

	myType := MyType{}

	fmt.Println(myType)  // I'm a MyType -- fmt.Println uses the Stringer interface as a parameter
}
```

### Interface Values

Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

`(value, type)`

An interface value holds a value of a specific underlying concrete type.

Calling a method on an interface value executes the method of the same name on its underlying type.

### Interface Values with Nil Underlying Values

If the concrete value inside the interface itself is `nil`, the method will be called with a `nil` receiver.

In some languages this would trigger a null pointer exception, but in Go it is common to write methods that gracefully handle being called with a `nil` receiver.

Note that an interface value that holds a `nil` concrete value is itself non-nil.

```aiignore
func (t *T) DoSomething() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println("not nil")
}
```

### Nil Interface Values

A nil interface value holds neither value nor concrete type.

Calling a method on a nil interface is a run-time error because there is no type inside the interface tuple to indicate which concrete method to call.

```aiignore
type I interface {
	M()
}

func main() {
	var i I  // (nil, nil)
	i.M()    // panic: runtime error: invalid memory address or nil pointer dereference
}
```

### The Empty Interface

The interface type that specifies zero methods is known as the empty interface:

`interface{}`

An empty interface may hold values of any type. (Every type implements at least zero methods.)

Empty interfaces are used by code that handles values of unknown type. For example, `fmt.Print` takes any number of arguments of type `interface{}`.

```aiignore
func main() {
	var i interface{}  // (nil, nil)

	i = 42  // (int, 42)

	i = "hello"  // (string, hello)
}
```

### Interface Composition

Golang prefers small interfaces that are built up with **composition** when functionality is needed.

```aiignore
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type ReaderWriter interface {
    Reader
    Writer
}
```

