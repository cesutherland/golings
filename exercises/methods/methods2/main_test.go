// maps3
// Make me compile!

// I AM NOT DONE
package main

import (
	"fmt"
	"testing"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func () IncrementAge() {  // something is missing
	p.Age++
}

func TestFullName(t *testing.T) {
	person := Person{"John", "Doe", 20}

	expectedFullName := "John Doe"

	fullName :=  // something is missing

	if fullName != expectedFullName {
		t.Errorf("full name should be %s but got %s", expectedFullName, fullName)
	}
}

func TestIncrementAge(t *testing.T) {
	person := Person{"John", "Doe", 20}

	expectedAge := 21

	person.IncrementAge()

	if person.Age != expectedAge {
		t.Errorf("age should be %d but got %d", expectedAge, person.Age)
	}
}
