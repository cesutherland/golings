// generics1
// Make me compile!

package main

import "fmt"

func main() {
	printMe("Hello, World!")
	printMe(42)
}

func printMe[V any](value V) {
	fmt.Println(value)
}
