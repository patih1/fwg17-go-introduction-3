package main

import (
	"fmt"
	"strings"
	"sync"
)

func wordByWord(sentence string, c chan string) {
	x := strings.Split(sentence, " ")
	for i := 0; i < len(x); i++ {
		c <- x[i]
	}
	close(c)
}

func cekPalindrom() {
	c := make(chan string)
	go wordByWord("sungguh malam yang masam", c)
	var wg sync.WaitGroup

	for result := range c {
		var reverse []byte

		for i := len(result) - 1; i >= 0; i-- {
			reverse = append(reverse, result[i])
		}

		wg.Add(1)
		result := result
		go func() {
			if string(reverse) == result {
				fmt.Println(result, " adalah palindrom")
			} else {
				fmt.Println(result, " bukan palindrom")
			}
			defer wg.Done()
		}()
	}
	wg.Wait()
}
