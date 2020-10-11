package sort

// Bubblesort function takes as argument a slice of integers and returns the sorted slice
// It uses the bubble-sort algorithm to sort the array
func Bubblesort(array []int) []int {
	returnval := make([]int, len(array))
	copy(returnval, array)
	for i := 0; i < len(returnval); i++ {
		for index := 0; index < len(returnval)-1; index++ {
			if returnval[index] > returnval[index+1] {
				swap(returnval, index, index+1)
			}
		}
	}
	return returnval
}

func swap(array []int, i int, j int) {
	temp := array[i]
	array[i] = array[j]
	array[j] = temp
}
