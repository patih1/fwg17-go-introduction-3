package main

import "fmt"

func deretBilangan(num int) []int {
	deretAngka := []int{}
	for i := 0; i <= num; i++ {
		deretAngka = append(deretAngka, i)
	}

	return deretAngka
}

type deretAngka struct {
	data []int
}

func (d deretAngka) prima() {
	var results []int = []int{}
	devider := 0
	for _, v := range d.data {
		devider = 0
		for i := 1; i <= v; i++ {
			if v%i == 0 {
				devider++
			}
		}
		if devider == 2 {
			results = append(results, v)
		}
	}
	fmt.Println("prima: ", results)
}

func (d deretAngka) ganjil() {
	var results []int = []int{}
	for _, v := range d.data {
		if v%2 != 0 {
			if v > 0 {
				results = append(results, v)
			}
		}
	}
	fmt.Println("ganjil: ", results)
}

func (d deretAngka) genap() {
	var results []int = []int{}
	for _, v := range d.data {
		if v%2 == 0 {
			if v > 0 {
				results = append(results, v)
			}
		}
	}
	fmt.Println("genap: ", results)
}

// [0, 1, 2, 3, 4, .... , 40]

func (d deretAngka) fibonacci() {
	nextPatern := 0
	var results []int = []int{}        // [0, 1, 1, 2, 3, 5, 8, 13]
	for i := 0; i < len(d.data); i++ { // i = 3
		if i <= 1 {
			results = append(results, d.data[i])
		} else if i > 1 {

			nextPatern = results[i-1] + results[i-2]

			if nextPatern > len(d.data)-1 {
				break
			}
			results = append(results, nextPatern)
		}
	}

	fmt.Println("fibonacci: ", results)
}
