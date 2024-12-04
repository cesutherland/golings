## Pointers

Go has pointers. A pointer holds the memory address of a value.

The type `*T` is a pointer to a `T` value. Its zero value is `nil`.

```aiignore
var p *int  // p is a pointer to an int
```

The `&` operator generates a pointer to its operand. 
It's helpful to think of this as the `address of` operator.

```
i := 42
p = &i  // p holds the address of i
```

The `*` operator denotes the pointer's underlying value. This is known as "dereferencing" or "indirecting".

```
fmt.Println(*p) // read i through the pointer p
*p = 21         // set i through the pointer p
```

Unlike C, Go has no pointer arithmetic (outside of the unsafe package).

Unlike C, Golang is a garbage collected language, so we do not need to manage memory manually.

### When to Use

Most of the time, we'll want to use a pointer when we want to share and modify data.

#### Common Uses of Pointers

* As mentioned above, functions/methods need to mutate data
* Some objects can't be copied (Mutex, WaitGroup)
* Some objects are too large to efficiently copy (> 64 bytes)
* When decoding protocol data into an object (JSON)
* To signal a `nil` object

### Function Parameters

Functions can accept pointers as parameters so that the underlying data can be changed.

```aiignore
func increment(i *int) {
	*i++
}

func main() {
	a := 1

	increment(&a)

	fmt.Println(a)  // 2
}
```
#### Reference Types

Note that certain types, slices, maps, and channels, are reference types that don't require
a pointer to be passed to functions in order to modify underlying data.

```aiignore
func modifySlice(s []int) {
	s[0] = 100
}

func main() {
	slice := []int{1, 2, 3}
	modifySlice(slice)
	fmt.Println(slice)  // [100 2 3]
}
```
### Pointers to Structs

#### Struct Fields

Struct fields can be accessed through a struct pointer.

To access the fields of a pointer to a struct, you can 
access them directly without an explicit dereference.

```aiignore
type Person struct {
	Name string
	Age  int
}

func main() {
	janePtr := &Person{"Jane", 25}

	janePtr.Age = 26 // access without explicit dereference

	fmt.Printf("%s is %d years old\n", janePtr.Name, janePtr.Age)  // Jane is 26 years old
}
```
