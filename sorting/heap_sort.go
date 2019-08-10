package sorting

import "../heap"

func HeapSort(arr []int) {
	h := heap.BuildHeapFromArray(arr, len(arr))
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		arr[i] = h.GetMax()
		h.ExtractMax()
	}
}
