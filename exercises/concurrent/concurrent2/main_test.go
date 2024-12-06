// concurrent2
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSendAndReceive(t *testing.T) {
	var buf bytes.Buffer

	messages := make(chan string)
	sendAndReceive(&buf, messages)

	got := buf.String()
	want := "HelloWorld"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func sendAndReceive(buf *bytes.Buffer, messages chan string) {
	go func() {
		messages <- "Hello"
		messages <- "World"
		close(messages)
	}()

	i := <-messages // here we just receive the first message from the channel; try the for range loop to get all messages
	fmt.Fprint(buf, i)
}
