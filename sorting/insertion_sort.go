package sorting

func InsertionSort(arr []int) {

	n := len(arr)

	if n < 2 {
		return
	}

	for i := 1; i < n; i++ {
		j := i - 1
		for j >= 0 && arr[i] < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		swap(&arr[j+1], &arr[i])
	}
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
