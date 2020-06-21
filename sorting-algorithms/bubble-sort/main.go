package main

import (
	"fmt"
)

// time complexity O(n ^ 2)
func bubbleSort(a []int, n int) {
	for k :=1; k < n-1; k++ {
		// if flag == 0, array is sorted because no swaps happened
		flag := 0
		
		// loop through array items checking if the first item is greater than the second
		// n-k-1 optimizes for greater than sorting since the end of the array is sorted
		// 1 cell further to the left every pass
		for i := 0; i <= n-k-1; i++ {
			if a[i] > a[i+1] {
				swap(a, i, flag)
			}
			flag = 1
		}
		fmt.Println("FLAG:", flag)
		if flag == 0 {
			continue
		} 
		fmt.Println(a)
	}
}

func swap(a []int, i int, flag int) {
	// replace next with current and current with next
	next := a[i]
	current := a[i+1]
	a[i] = current
	a[i+1] = next			
}

func main() {
	a := []int{2,4,1,5,7,3}
	fmt.Println(a)
	fmt.Println(len(a))
	bubbleSort(a, len(a))
}