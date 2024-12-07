// loops1
// Make me compile!

package main

import "fmt"

func main() {
	// normal for loop
	for i := 0; i < 11; i++ { // let's print to ten
		fmt.Println(i)
	}

	// while loop
	num := 1

	for num < 11 { // let's print to 10 using a `while` loop
		fmt.Println(num)
		num++
	}
}
