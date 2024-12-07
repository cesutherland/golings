// anonymous functions2
// Make me compile!
package main

import "fmt"

func main() {
	var sayBye func(name string)

	sayBye = func(name string) {
		fmt.Printf("Bye %s", name)
	}

	sayBye("Carl")
}
