package sort_test

import (
	"testing"

	"github.com/shashwat183/gosort/sort"
)

// test for selection sort
func TestSelectionsort(t *testing.T) {

	// test for long array with no repeated numbers
	array := []int{1, 4, 5, 3, 6, 9}
	sorted := []int{1, 3, 4, 5, 6, 9}
	got := sort.Selectionsort(array)

	if !compareslices(got, sorted) {
		t.Errorf("array not sorted correctly %v", got)
	} else {
		t.Logf("array sorted successfully %v", got)
	}

	// test for long array with repeated elements
	array = []int{1, 4, 4, 5, 3, 6, 5, 9}
	sorted = []int{1, 3, 4, 4, 5, 5, 6, 9}
	got = sort.Selectionsort(array)

	if !compareslices(got, sorted) {
		t.Errorf("array not sorted correctly %v", got)
	} else {
		t.Logf("array sorted successfully %v", got)
	}

	// test for array with 1 element
	array = []int{4}
	sorted = []int{4}
	got = sort.Selectionsort(array)

	if !compareslices(got, sorted) {
		t.Errorf("array not sorted correctly %v", got)
	} else {
		t.Logf("array sorted successfully %v", got)
	}
}

// test for bubble sort
func TestBubblesort(t *testing.T) {

	// test for long array with no repeated numbers
	array := []int{1, 4, 5, 3, 6, 9}
	sorted := []int{1, 3, 4, 5, 6, 9}
	got := sort.Bubblesort(array)

	if !compareslices(got, sorted) {
		t.Errorf("array not sorted correctly %v", got)
	} else {
		t.Logf("array sorted successfully %v", got)
	}

	// test for long array with repeated elements
	array = []int{1, 4, 4, 5, 3, 6, 5, 9}
	sorted = []int{1, 3, 4, 4, 5, 5, 6, 9}
	got = sort.Bubblesort(array)

	if !compareslices(got, sorted) {
		t.Errorf("array not sorted correctly %v", got)
	} else {
		t.Logf("array sorted successfully %v", got)
	}

	// test for array with 1 element
	array = []int{4}
	sorted = []int{4}
	got = sort.Bubblesort(array)

	if !compareslices(got, sorted) {
		t.Errorf("array not sorted correctly %v", got)
	} else {
		t.Logf("array sorted successfully %v", got)
	}
}

// test for merge sort
func TestMergesort(t *testing.T) {

	// test for long array with no repeated numbers
	array := []int{1, 4, 5, 3, 6, 9}
	sorted := []int{1, 3, 4, 5, 6, 9}
	got := sort.Mergesort(array)

	if !compareslices(got, sorted) {
		t.Errorf("array not sorted correctly %v", got)
	} else {
		t.Logf("array sorted successfully %v", got)
	}

	// test for long array with repeated elements
	array = []int{1, 4, 4, 5, 3, 6, 5, 9}
	sorted = []int{1, 3, 4, 4, 5, 5, 6, 9}
	got = sort.Mergesort(array)

	if !compareslices(got, sorted) {
		t.Errorf("array not sorted correctly %v", got)
	} else {
		t.Logf("array sorted successfully %v", got)
	}

	// test for array with 1 element
	array = []int{4}
	sorted = []int{4}
	got = sort.Mergesort(array)

	if !compareslices(got, sorted) {
		t.Errorf("array not sorted correctly %v", got)
	} else {
		t.Logf("array sorted successfully %v", got)
	}

	// test for array with descending sort
	array = []int{5, 4, 3, 2, 1}
	sorted = []int{1, 2, 3, 4, 5}
	got = sort.Mergesort(array)

	if !compareslices(got, sorted) {
		t.Errorf("array not sorted correctly %v", got)
	} else {
		t.Logf("array sorted successfully %v", got)
	}
}

func compareslices(a, b []int) bool {
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
