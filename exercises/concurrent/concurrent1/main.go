// concurrent1
// Make me compile!

package main

import "fmt"

func main() {
	messages := make(chan string) // we need to define a chan of a type, what is that type here?

	go func() { messages <- "ping" }() // what do we need to add to run this in a `goroutine`

	msg := <-messages

	fmt.Println(msg)
}
