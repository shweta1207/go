package heap

// Heap is a complete binary tree
// Nodes[0...n-1]
// Root Node :0
// left child : 2n+1
// right child : 2n+2
// parent : floor(n-1/2)

func (h *Heap) GetMax() int {
	if h.Size() == 0 {
		return 0
	}
	return h.heap[0]
}

// Assumes both its children are min heaps
func (h *Heap) MaxHeapify(index int) {
	n := h.Size()
	if n == 0 || n <= index {
		return
	}
	left := 2*index + 1
	right := 2*index + 2
	largest := index
	if left < n && h.heap[index] < h.heap[left] {
		largest = left
	}
	if right < n && h.heap[largest] < h.heap[right] {
		largest = right
	}
	if largest != index {
		swap(&h.heap[index], &h.heap[largest])
		h.MaxHeapify(largest)
	}
}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func (h *Heap) ExtractMax() {
	n := h.Size()
	if n == 0 {
		return
	}
	h.heap[0] = h.heap[n-1]
	h.heap = h.heap[:n-1]
	h.MaxHeapify(0)
}
func (h *Heap) InsertMax(val int) {
	h.heap = append(h.heap, val)
	n := h.Size()
	curr := n - 1
	par := (curr - 1) / 2
	for par >= 0 && h.heap[par] < h.heap[curr] {
		swap(&h.heap[par], &h.heap[curr])
		curr = par
		par = (curr - 1) / 2
	}
}

// Decrease value of key at index ind to val
func (h *Heap) IncreaseKey(ind int, val int) {
	n := h.Size()
	if n <= ind {
		return
	}
	if ind == 0 {
		h.heap[ind] = val
	}
	par := (ind - 1) / 2
	for par >= 0 && h.heap[par] < h.heap[ind] {
		swap(&h.heap[par], &h.heap[ind])
		ind = par
		par = (ind - 1) / 2
	}
}

// Delete key at given index
func (h *Heap) DeleteMax(ind int) {
	n := h.Size()
	if n <= ind {
		return
	}
	h.DecreaseKey(ind, 999) // INF
	h.ExtractMax()
}
