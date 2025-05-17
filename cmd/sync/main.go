package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	counter := 0

	// Function to increment the counter with mutex lock
	increment := func(id int, wg *sync.WaitGroup) {
		defer wg.Done() // Decrement the counter when the goroutine completes
		fmt.Printf("Goroutine %d: Waiting to lock\n", id)
		mu.Lock() // Lock the mutex
		fmt.Printf("Goroutine %d: Locked\n", id)

		// Critical section
		counter++
		fmt.Printf("Goroutine %d: Counter = %d\n", id, counter)

		time.Sleep(1 * time.Second) // Simulate some work
		mu.Unlock()                 // Unlock the mutex
		fmt.Printf("Goroutine %d: Unlocked\n", id)
	}

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go increment(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Final Counter:", counter)
}
