package main

import "fmt"

// Pre-requisite is recursion
func mergeSort(a[]int, n int) {
	if n < 2 {
		return
	}

	mid := len(a)/2
	left := make([]int, mid)
	right := make([]int, n-mid)

	for i := 0; i <= mid -1; i++ {
		left[i] = a[i]
	}
	for i := mid; i <= n-1; i++ {
		right[i-mid] = a[i]
	}
	mergeSort(left, len(left))
	mergeSort(right, len(right))
	merge(left,right,a)
}

func merge(left []int, right []int, a []int) {
	var i,j,k int
	// deals with both left and right arrays
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			a[k] = left[i]
			i++
		} else {
			a[k] = right[j]
			j++
		}
		k++
	}
	// one of arrays is exhausted
	// if left array still has unpicked values
	for i < len(left) {
		a[k] = left[i]
		i++
		k++
	}
	// if right array still has unpicked values
	for j < len(right) {
		a[k] = right[j]
		j++
		k++
	}
}

func main() {
	a := []int{2,4,1,6,8,5,3,7}
	fmt.Println("Unsorted: ", a)
	fmt.Println("length: ", len(a))
	mergeSort(a, len(a))
	fmt.Println("Sorted: ", a)
}

// time complexity O(n log n)