// loops2
// Make me compile!
package main

import "fmt"

func main() {

	evenNumbers := []int{2, 4, 6, 8, 10}

	for _, v := {  // print all the evenNumbers using range
		fmt.Printf("%d is even\n", v)
	}

	phoneBook := map[string]string{
		"Ana":  "+01 101 102",
		"John": "+01 333 666",
	}

	for phone, name := range phoneBook {
		fmt.Printf("%s has the %s phone\n", name, phone)
	}
}
