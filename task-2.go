package main

import (
	"fmt"
	"sync"
)

func fibonacci(chn chan int) {

	x, y := 0, 1

	for i := 0; i < 10; i++ {
		chn <- x
		x, y = y, x+y
	}
	// mengindikasikan tidak ada value diterima dan channel ditutup
	close(chn)
}

func ganjilGenap() {
	var wg sync.WaitGroup

	chn := make(chan int)
	go fibonacci(chn)

	for result := range chn {
		wg.Add(1)
		result := result
		go func() {
			if result%2 == 0 {
				fmt.Println("genap ", result)
			} else {
				fmt.Println("ganjil ", result)
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
}
