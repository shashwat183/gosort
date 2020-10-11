package sort

// Mergesort function takes as argument a slice of integers and returns the sorted slice
// It uses the marge-sort algorithm to sort the array
func Mergesort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	middle := len(slice) / 2
	left := slice[:middle]
	right := slice[middle:]

	sortedleft := Mergesort(left)
	sortedright := Mergesort(right)

	sorted := merge(sortedleft, sortedright)
	return sorted
}

func merge(left []int, right []int) []int {
	merged := make([]int, len(left)+len(right))
	i := 0
	j := 0
	k := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged[k] = left[i]
			i++
		} else {
			merged[k] = right[j]
			j++
		}
		k++
	}

	for i < len(left) {
		merged[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		merged[k] = right[j]
		j++
		k++
	}

	return merged
}
