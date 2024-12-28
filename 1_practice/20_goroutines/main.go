package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	
    for i := 0; i < 5; i++ {
        fmt.Println(msg)
        time.Sleep(500 * time.Millisecond)
    }
}

func main() {
	var wg sync.WaitGroup // Create a WaitGroup

    wg.Add(2) // Add 2 goroutines to the WaitGroup
   

    go printMessage("Goroutine 1", &wg)
    go printMessage("Goroutine 2", &wg)

    wg.Wait() // Wait for all goroutines to finish
    fmt.Println("All goroutines finished")


}
