package sort

import "fmt"

// Quicksort function takes as argument a slice of integers and returns the sorted slice
// It uses the bubble-sort algorithm to sort the array
func Quicksort(slice []int) []int {
	returnval := make([]int, len(slice))
	copy(returnval, slice)
	sort(returnval, 0, len(returnval)-1)
	return returnval
}

func sort(slice []int, low int, high int) {
	if low < high {
		partitionindex := partition(slice, low, high)

		sort(slice, low, partitionindex-1)
		sort(slice, partitionindex+1, high)
	}
}

func partition(slice []int, low int, high int) int {
	i := low - 1
	pivot := slice[high]

	for j := low; j < high; j++ {

		if slice[j] < pivot {
			i++
			fmt.Println(i)
			swap(slice, i, j)
		}
	}

	swap(slice, i+1, high)

	return i + 1
}
