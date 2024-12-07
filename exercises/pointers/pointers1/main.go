// pointers1
// Make me compile!

package main

import "fmt"

func swap(a, b *int) {
	a, b = b, a
}

func main() {

	x := 10 // do not change
	y := 20 // do not change

	swap(&x, &y) // do not change

	// Fix the code so it compiles and the output is "x is 20 and y is 10"
	fmt.Printf("x is %d and y is %d\n", x, y)
}
