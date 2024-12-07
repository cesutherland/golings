// methods1
// Make me compile!

package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) UpdateName(name string) Person {
	p.Name = name

	return p
}

func main() {
	person := Person{"John", 25}

	// update person's name to be Jane using the method
	person.UpdateName("Jane")

	fmt.Println("Name is", person.Name)
}
