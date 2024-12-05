// generics2
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

type Number interface {
	int // ensure we cover both types in the call to addNumbers below
}

func main() {
	fmt.Println(addNumbers(1, 2))
	fmt.Println(addNumbers(1.0, 2.3))
}

func addNumbers[](n1, n2 T) T { // a type needs to be defined
	return n1 + n2
}
