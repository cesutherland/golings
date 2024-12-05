// pointers2
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func UpdateName(p *Person, name string) {
	p.Name = name
}

func main() {

	person := Person{"John", 42} // do not change this line

	// update person's name to Jane using the UpdateName function above
	UpdateName()

	// Fix the code so it compiles and the output is "person = Jane"
	fmt.Printf("person = %s\n", person.Name)
}
