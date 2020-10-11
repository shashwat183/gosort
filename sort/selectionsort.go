package sort

// Selectionsort function takes as input an int slice and returns the the sorted slice.
// It uses the quick-sort algorithm to sort the array
func Selectionsort(array []int) []int {
	if len(array) == 1 {
		return array
	}
	index := 0
	smallest := array[index]
	for i, num := range array {
		if num < smallest {
			smallest = num
			index = i
		}
	}
	var sortarray []int = []int{smallest}
	cutarray := cutarray(array, index)
	sortarray = append(sortarray, Selectionsort(cutarray)...)
	return sortarray
}

func cutarray(slice []int, index int) []int {
	slice[index] = slice[len(slice)-1]
	slice[len(slice)-1] = 0
	slice = slice[:len(slice)-1]
	return slice
}
