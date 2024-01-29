package main

import (
	"time"
)

func main() {
	// try-it!
	// a := []int{7, 10, 2, 34, 33, -12, -8, 4}

	// chn := make(chan int)

	// go sum(a[:len(a)/2], chn) // 53
	// go sum(a[len(a)/2:], chn) // 17

	// x, y := <-chn, <-chn

	// fmt.Println(x, " ", y)

	// time.Sleep(time.Second)
	//=================================================//

	//task-1
	deret := deretBilangan(40)
	var data deretAngka = deretAngka{
		deret,
	}

	go func() {
		data.prima()
		data.ganjil()
		data.genap()
		data.fibonacci()
	}()
	time.Sleep(time.Second)

	//task-2
	ganjilGenap()

	//task-3
	cekPalindrom()
}
