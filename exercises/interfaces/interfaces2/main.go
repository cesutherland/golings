// interfaces2
// Make me compile!

// I AM NOT DONE
package main

import "fmt"

func describe(e) string { // which interface can be satisfied by any type?
	return fmt.Sprintf("(%T, %[1]v)", e)
}

func main() {
	i := 10
	s := "Hello"
	o := struct {
		name string
		age  int
	}{"John", 20}

	fmt.Println(describe(i))
	fmt.Println(describe(s))
	fmt.Println(describe(o))
}
