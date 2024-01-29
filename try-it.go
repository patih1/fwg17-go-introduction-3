package main

func sum(d []int, ch chan int) {
	sum := 0
	for _, v := range d {
		sum += v
	}
	ch <- sum
}
