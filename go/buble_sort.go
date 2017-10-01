package main

import "fmt"

func main() {
	a := []int{10, 44, 32, 56, 73, 11, 1, 23, 2}
	fmt.Println(a)

	bubble(a)
	fmt.Println(a)
}

func bubble(a []int) {
	n := len(a)

	for n > 1 {
		for i := 0; i < n-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		n--
	}
}
