package sorting

func MergeSort(arr []int, l int, r int) {

	if l < r {
		m := l + (r-l)/2
		MergeSort(arr, l, m)
		MergeSort(arr, m+1, r)
		Merge(arr, l, m, r)
	}

}

func Merge(arr []int, l, m, r int) {
	n1 := m - l + 1
	n2 := r - m
	larr := make([]int, n1)
	rarr := make([]int, n2)

	for i := 0; i < n1; i++ {
		larr[i] = arr[l+i]
	}
	for i := 0; i < n2; i++ {
		rarr[i] = arr[m+1+i]
	}
	i, j, k := 0, 0, 0
	for i < n1 && j < n2 {
		if larr[i] < rarr[j] {
			arr[k] = larr[i]
			i++
			k++
		} else {
			arr[k] = rarr[j]
			j++
			k++
		}

	}
	for i < n1 {
		arr[k] = larr[i]
		i++
		k++
	}
	for i < n1 {
		arr[k] = rarr[j]
		j++
		k++
	}
}
