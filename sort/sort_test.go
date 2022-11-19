package sort

import (
	"sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	values := []int{1, 8, 3, 9, 4, 5, 7}
	QuickSort(&values)

	if !sort.IntsAreSorted(values) {
		t.Errorf("quicksort did not sort the values")
	}
}
