// variables4
// Make me compile!

package main

import "fmt"

func main() {
	x := "TEN" // Don't change this line
	fmt.Printf("x has the value %s\n", x)

	if true {
		x = 1 // make this a shadowed local copy of x
		fmt.Println(x + 1)
	}

	fmt.Println(x)
}
