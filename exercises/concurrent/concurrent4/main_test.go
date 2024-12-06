// concurrent4
// Make the tests pass!

// I AM NOT DONE
package main_test

import (
	"sync"
	"testing"
)

type SafeCounter struct {
	mu    sync.Mutex
	value int
}

func (c *SafeCounter) Inc() {
	c.value++ // many goroutines trying to update the counter? maybe something in the SafeCounter can help!
}

func (c *SafeCounter) Value() int {
	return c.value // many goroutines trying to update the counter? maybe something in the SafeCounter can help!
}

func TestSafeCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := SafeCounter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		if counter.Value() != 3 {
			t.Errorf("got %d, want 3", counter.Value())
		}
	})

	t.Run("concurrent increments", func(t *testing.T) {
		counter := SafeCounter{}
		numIncrements := 1000

		var wg sync.WaitGroup
		wg.Add(numIncrements)

		for i := 0; i < numIncrements; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		if counter.Value() != numIncrements {
			t.Errorf("got %d, want %d", counter.Value(), numIncrements)
		}
	})
}
