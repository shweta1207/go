package sorting

// SelectionSort sorts by placing smallest element at the start of the unsorted array
func SelectionSort(arr []int) {

	n := len(arr)
	for i := 0; i < n-1; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		swap(&arr[i], &arr[min])
	}
}
