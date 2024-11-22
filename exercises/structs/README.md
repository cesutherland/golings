## Structs

Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

```aiignore
type Person struct {
    Name string
    Age  int
}
```

### Initialization

Initializing a struct can take several forms.

```aiignore
person := Person{
    "John",
    25,
}

fmt.Println(person) // {John 25}

person = Person{
    Name: "Jane",  // you can name the fields when initializing a struct
    Age:  25,
}

fmt.Println(person) // {Jane 25}
```

Omitted fields will default to the zero value of the field type.

```aiignore
person := Person{}
fmt.Println(person) // { 0}

person = Person{Name: "John"}
fmt.Println(person) // {John 0}
```

It’s idiomatic to encapsulate new struct creation in constructor functions.

```aiignore
func NewPerson(name string, age int) Person {
	return Person{name, age}
}

func main {

    person := NewPerson("John", 25)
    fmt.Println(person) // {John 25}
}
```

### Field Access

Struct fields (value or pointer structs) can be accessed through dot notation.

```aiignore
person := Person{"John", 25}
fmt.Println(person.Name) // John
```

### Mutation

Structs are mutable

```aiignore
person := Person{"John", 25}
person.Age = 30
fmt.Println(person) // {John, 30}
```

### Anonymous Structs

If a struct type is only used for a single value, we don’t have to give it a name. 
The value can have an anonymous struct type.

```aiignore
person := struct {
		name string
		age  int
	}{
		"John",
		25,
	}
```

### Empty Struct

An empty struct is a struct with no fields that takes up zero memory. They behave just like any other struct.
There are several use cases for them, but you'll commonly see them use in conjunction with channels as a signal or
when building a `set` as Golang does not have this type.

```aiignore
// struct{}  // empty struct

// Set

set := make(map[string]struct{})

set["John"] = struct{}{} // an instance of an empty struct
set["Jane"] = struct{}{}
set["John"] = struct{}{}

if _, exists := set["Jane"]; exists {
    fmt.Println("Jane is in the set")
}


```

### Struct Embedding


Go supports embedding of structs and interfaces to express a more seamless composition of types.

```aiignore
type Address struct {
	City  string
	State string
}

type Person struct {
	Address          // embedded struct
	Name    string
	Age     int
}

func (a Address) PrintAddress() string {
	return fmt.Sprintf("State: %v, City: %v", a.State, a.City)
}

person := Person{
		Name: "John",
		Age:  25,
		Address: Address{
			City:  "New York",
			State: "New York",
		},
	}

	fmt.Println(person)                // {{New York New York} John 25}
	fmt.Println(person.City)           // New York - City is "promoted"
	fmt.Println(person.Address.City)   // New York - City is accessed through the full path's type name
	fmt.Println(person.PrintAddress()) // State: New York, City: New York - PrintAddress is accessed directly from Person
```

#### Field and Method Promotion

As shown above, an embedded structs fields and methods are "promoted"
and accessible on the outer struct directly.

`person.City` will access the Person.Address.City field directly.

`person.PrintAddress()` will access Address' promoted PrintAddress method.

### Comparable

Structs are comparable if all it's fields are comparable

```aiignore
person1 := Person{
		Name: "John",
		Age:  25,
	}

	person2 := Person{
		Name: "John",
		Age:  25,
	}

	person3 := Person{
		Name: "Jane",
		Age:  25,
	}

	fmt.Println(person1 == person2)  // true
	fmt.Println(person1 == person3)  // false
```

### Passed By

Structs are passed by value.