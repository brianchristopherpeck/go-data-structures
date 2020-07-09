package main

import "fmt"

// most built-in language sort functions are quick sort implementations
// quickSort divides the array based on a pivot point
func quickSort(a []int, start int, end int) {
	if start >= end {
		return
	}
	partitionIndex := partition(a, start, end - 1)
	quickSort(a, start, partitionIndex - 1)
	quickSort(a, partitionIndex + 1, end)
}

func partition(a []int, start int, end int) int {
	pivot := a[end]
	partitionIndex := start
	for i := start; i <= end; i++ {
		if a[i] < pivot {
			swap(a, i, partitionIndex)
			partitionIndex++
		}
	}
	swap(a, partitionIndex, end)

	return partitionIndex
}

func swap(a []int, j int, i int) {
	first := a[j]
	second := a[i]
	a[j] = second
	a[i] = first
}

func main() {
	a := []int{7,2,1,6,8,5,3,4}
	fmt.Println(a)
	quickSort(a, 0, len(a))
	fmt.Println(a)
}