package sort

// QuickSort sorts the provided values using the QuickSort algorithm.
func QuickSort(values *[]int) {
	quickSort(values, 0, len(*values)-1)
}

// The internal quicksort function that can operate on subarrays of the values.
func quickSort(values *[]int, low int, high int) {
	if low < high {
		partitionIdx := partition(values, low, high)
		quickSort(values, low, partitionIdx-1)
		quickSort(values, partitionIdx+1, high)
	}
}

// Rearranges `values` such that we first have the elements lower than the pivot, then the pivot,
// then the pivot, then the elements larger than the pivot. The pivot is chosen as the element of
// `values` at index `high`.
func partition(values *[]int, low int, high int) int {
	pivot := (*values)[high]
	firstUnsorted := low - 1

	for idx := low; idx < high; idx++ {
		if (*values)[idx] < pivot {
			firstUnsorted++
			swap(values, firstUnsorted, idx)
		}
	}

	swap(values, firstUnsorted+1, high)

	return firstUnsorted + 1
}

// Swaps the elements of `values` at indices `i` and `j`.
func swap(values *[]int, i int, j int) {
	temp := (*values)[i]
	(*values)[i] = (*values)[j]
	(*values)[j] = temp
}
