// generics1
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

func main() {
	printMe("Hello, World!")
	printMe(42)
}

func printMe[](value) { // we're missing a type parameter here that would allow both string and int or maybe any
	fmt.Println(value)
}
