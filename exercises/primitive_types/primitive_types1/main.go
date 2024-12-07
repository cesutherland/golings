// primitive_types1
// Make me compile!

package main

import "fmt"

func main() {
	storeIsOpen := true
	if storeIsOpen {
		fmt.Println("The store is open, let's buy some clothes!")
	}

	// update the value to make sure this prints
	storeIsOpen = false
	if !storeIsOpen {
		fmt.Println("Oh no, let's buy some clothes online!")
	}
}
