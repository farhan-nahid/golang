package main

import (
	"fmt"
	"sync"
)

// Counter struct with a mutex
type Counter struct {
	mu    sync.Mutex
	value int
}

// Increment safely increments the counter value
func (c *Counter) Increment() {
	c.mu.Lock()   // Acquire the lock
	c.value++     // Modify the shared resource
	c.mu.Unlock() // Release the lock
}

// Value safely retrieves the counter value
func (c *Counter) Value() int {
	c.mu.Lock()         // Acquire the lock
	defer c.mu.Unlock() // Ensure lock is released
	return c.value
}

func main() {
	var wg sync.WaitGroup
	counter := Counter{}

	// Increment counter concurrently
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Print the final counter value
	fmt.Println("Final Counter Value:", counter.Value())
}
