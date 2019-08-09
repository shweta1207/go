package binarysearch

// BinarySearch return the index of the element to be searched
func BinarySearch(arr []int, l int, r int, x int) int {

	if l <= r {

		m := l + (r-l)/2

		if arr[m] == x {
			return m
		}

		if arr[m] < x {
			return BinarySearch(arr, l, m-1, x)
		}

		return BinarySearch(arr, m+1, r, x)
	}
	return -1

}
