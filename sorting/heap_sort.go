package sorting

import "github.com/shweta1207/go/heap"

func HeapSort(arr []int) {
	h := heap.BuildHeapFromArray(arr, len(arr))
	n := len(arr)
	for i := n - 1; i >= 0; i-- {
		arr[i] = h.GetMax()
		h.ExtractMax()
	}
}
