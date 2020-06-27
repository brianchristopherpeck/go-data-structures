package main

import "fmt"

func insertionSort(a []int, n int) {
	for i := 1; i < n; i++ {
		value := a[i]
		hole := i
		for hole > 0 && a[hole -1] > value {
			a[hole] = a[hole - 1]
			hole = hole - 1
		}
		a[hole] = value
		fmt.Println(a)
	}
}

func main() {
	a := []int{2,7,4,1,5,3}
	insertionSort(a, len(a))
	fmt.Println(a)
}

// 2,7,4,1,5,3
// 2,4,7,1,5,3
// 1,2,4,7,5,3
// 1,2,4,5,7,3
// 1,2,3,4,5,7

