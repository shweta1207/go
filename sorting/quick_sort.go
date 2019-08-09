package sorting

func QuickSort(arr []int, l int, r int) {

	if l < r {
		p := Partition(arr, l, r)
		QuickSort(arr, l, p-1)
		QuickSort(arr, p+1, r)
	}

}
func Partition(arr []int, l int, r int) int {
	pivot := arr[r]
	i := l - 1
	for k := l; k < r-1; k++ {
		if arr[k] < pivot {
			i++
			swap(&arr[i], &arr[k])
		}
	}
	swap(&arr[i+1], &arr[r])
	return i + 1
}
