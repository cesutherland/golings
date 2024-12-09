# Concurrency

## Goroutines

A goroutine is a lightweight thread managed by the Go runtime.

`go f(x, y, z)`

starts a new goroutine running

`f(x, y, z)`

The evaluation of `f`, `x`, `y`, and `z` happens in the current goroutine and the execution of `f` happens in 
the new goroutine.

```aiignore
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
```

Goroutines run in the same address space, so access to shared memory must be synchronized. The sync package provides 
useful primitives, although you won't need them much in Go as there are other primitives.

## Channels

Channels are a typed conduit through which you can send and receive values with the channel operator, `<-`.

```
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and assign value to v.
```
**The data flows in the direction of the arrow.**

Like maps and slices, channels must be created before use:

`ch := make(chan int)`

**By default, sends and receives block until the other side is ready.** This allows goroutines to synchronize without 
explicit locks or condition variables.

The example code below sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines 
have completed their computation, it calculates the final result.

```aiignore
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum  // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c  // receive from c

	fmt.Println(x, y, x+y)
}
```

## Buffered Channels

Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel.

`ch := make(chan int, 100)`

**Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.**

```aiignore
func main() {
	ch := make(chan int, 2)  // buffer len of 2
	ch <- 1
	ch <- 2
	ch <- 3  // try to send a 3rd
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)  // fatal error: all goroutines are asleep - deadlock!
}
```

## Range and Close

A sender can `close` a channel to indicate that no more values will be sent. Receivers can test whether a channel has 
been closed by assigning a second parameter to the receive expression.

`v, ok := <-ch`

`ok` is `false` if there are no more values to receive and the channel is closed.

The loop for `i := range c` receives values from the channel repeatedly until it is closed.

**Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.**

**Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the 
receiver must be told there are no more values coming, such as to terminate a range loop.**

```aiignore
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
```

## Select

The `select` statement lets a goroutine wait on multiple communication operations.

A `select` blocks until one of its cases can run, then it executes that case. **It chooses one at random if multiple 
are ready.**

```aiignore
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
```

### Default Selection

The `default` case in a `select` is run if no other case is ready.

Use a `default` case to try a send or receive without blocking:

```
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

## Mutex

We've seen how channels are great for communication among goroutines.

But what if we don't need communication? What if we just want to make sure only one goroutine can access a variable 
at a time to avoid conflicts?

This concept is called _mutual exclusion_, and the conventional name for the data structure that provides it is `mutex`.

Go's standard library provides mutual exclusion with `sync.Mutex` and its two methods:

```
Lock`
Unlock
```

We can define a block of code to be executed in mutual exclusion by surrounding it with a call to `Lock` and `Unlock` 
as shown on the `Inc` method below.

We can also use `defer` to ensure the `mutex` will be unlocked as in the `Value` method.

```aiignore
// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```

## WaitGroups

A `WaitGroup` waits for a collection of goroutines to finish. The main goroutine calls `WaitGroup.Add` to set the 
number of goroutines to wait for. Then each of the goroutines runs and calls `WaitGroup.Done` when finished. 
At the same time, `WaitGroup.Wait` can be used to block until all goroutines have finished.

**A WaitGroup must not be copied after first use.**

```aiignore
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()  // Decrement the counter when the goroutine finishes

	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Launch 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1)  // Increment the counter for each goroutine
		go worker(i, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("All workers done")
}
```