// concurrent3
// Make the tests pass!

package main_test

import (
	"testing"
	"time"
)

func TestSelectChannel(t *testing.T) {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		c1 <- "from c1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		c2 <- "from c2"
	}()

	select {
	case msg1 := <-c1: // make sure to receive from the right channel
		if msg1 != "from c1" {
			t.Errorf("Expected 'from c1', got %s", msg1)
		}
	case msg2 := <-c2: // make sure to receive from the right channel
		if msg2 != "from c2" {
			t.Errorf("Expected 'from c2', got %s", msg2)
		}
	}
}
