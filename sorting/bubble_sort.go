package sorting

// BubbleSort sorts by repeatedly swapping adjacent elements if a[i] > a[i+1]
func BubbleSort(arr []int) {

	n := len(arr)

	if n < 2 {
		return
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(&arr[j], &arr[j+1])
			}
		}
	}
}
