package main

import "fmt"

func selectionSort(a []int, n int) {
	// loop through the array
	for i := 0; i < n -1; i++ {
		// set the index to save the smallest value
		var imin = i

		// inner loop
		for j := i + 1; j < n; j++ {
			// if a[j] is less than current minimum, that becomes the current minumum
			if a[j] < a[imin] {
				imin = j
			}
		}
		// save current value for swapping
		var currentValue = a[i]

		// swap with current value with current minimum
		a[i] = a[imin]

		// swap current minimum with current value
		a[imin] = currentValue
		fmt.Println(a)
	}
}

func main() {
	a := []int{2,7,4,1,5,3}
	selectionSort(a, len(a))
	fmt.Println(a)
}

// 2,7,4,1,5,3
// 1,7,4,2,5,3
// 1,2,4,7,5,3
// 1,2,3,7,5,4
// 1,2,3,4,5,7

// time complexity O(n^2)