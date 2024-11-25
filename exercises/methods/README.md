## Methods

Go does not have classes. However, you can define methods on types.

A method is a function with a special receiver argument.

The receiver appears in its own argument list between the `func` keyword and the method name.

Below, the `Person` struct has a method `UpdateName` that is used to set the name.
```aiignore
type Person struct {
	Name string
}

func (p Person) UpdateName(name string) Person {
	p.Name = name

	return p
}

func main() {
	person := Person{"John"}

	person = person.UpdateName("Jane")

	fmt.Println(person.Name)  // Jane
}
```

**A method is just a function with a special receiver.**

The above example could be rewritten as a function like so:

```aiignore
type Person struct {
	Name string
}

func UpdateName(p Person, name string) Person {
	p.Name = name

	return p
}

func main() {
	person := Person{"John"}

	person = UpdateName(person, "Jane")

	fmt.Println(person.Name)
}
```

You can declare a method on non-struct types, too.

```aiignore
type MyInt int64

func (m MyInt) Increment() MyInt {
	m++
	return m
}

func main() {
	myInt := MyInt(1)

	nextNumber := myInt.Increment()

	fmt.Println(nextNumber)  // 2
}
```

### Pointer Receivers

You can declare methods with pointer receivers as well.

This means the receiver type has the literal syntax `*T` for some type `T`.

Methods with pointer receivers can modify the value to which the receiver points. 
Since methods often need to modify their receiver, pointer receivers are more common than value receivers.

Below, the `UpdateName` method has a pointer receiver (`p *Person`). Now, the
method can directly modify the `Person` receiver without needing to return it.
```aiignore
func (p *Person) UpdateName(name string) {
	p.Name = name
}

func main() {
	personPtr := &Person{"John"}

	personPtr.UpdateName("Jane")

	fmt.Println(personPtr.Name) // Jane
}
```

**Note that in many cases (when the value is addressable) a pointer method can be invoked on a value. 
The Go compiler will automatically add the & operator**

For example, you could also invoke the method on a value:

```aiignore
person := Person{"John"}

person.UpdateName("Jane")
```

### Choosing Pointer vs Value Receivers

There are two reasons to use a pointer receiver.

The first is so that the method can modify the value that its receiver points to.

The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

**In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both.**
