package main

import (
	"time"
	"sync"
	"fmt"
)

func main() {
	begin := time.Now()

	wg := &sync.WaitGroup{}

	// add items to my waitgroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Printf("starting %d...\n", i)
			time.Sleep(1 * time.Second)
			fmt.Printf("ending %d...\n", i)
		}(i)

	}

	fmt.Println("waiting...")
	wg.Wait()

	// print off the duration the program took to run
	fmt.Println(time.Since(begin))
}
