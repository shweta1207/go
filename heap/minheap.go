package heap

// Heap is a complete binary tree
// Nodes[0...n-1]
// Root Node :0
// left child : 2n+1
// right child : 2n+2
// parent : floor(n-1/2)

type Heap struct {
	heap []int
}

func NewHeap(cap int) *Heap {
	return &Heap{
		heap: make([]int, 0, cap),
	}
}

func (h *Heap) Size() int {
	return len(h.heap)
}

func (h *Heap) GetMin() int {
	if h.Size() == 0 {
		return 0
	}
	return h.heap[0]
}

func BuildHeapFromArray(arr []int, cap int) *Heap {
	h:=&Heap{
		heap:make([]int,len(arr),cap),
	}
	n:=len(arr)
	h[:n]=arr[:n]
	for i:=n/2-1;i>=0;i--{
		h.MinHeapify(i)
	}
}

// Assumes both its children are min heaps
func (h *Heap) MinHeapify(index int) {
	n := h.Size()
	if n == 0 || n <= index {
		return
	}
	left := 2*index + 1
	right := 2*index + 2
	smallest := index
	if left < n && h.heap[index] > h.heap[left] {
		smallest = left
	}
	if right < n && h.heap[smallest] > h.heap[right] {
		smallest = right
	}
	if smallest != index {
		swap(&h.heap[index], &h.heap[smallest])
		h.MinHeapify(smallest)
	}
}


func (h *Heap) ExtractMin() {
	n := h.Size()
	if n == 0 {
		return
	}
	h.heap[0] = h.heap[n-1]
	h.heap = h.heap[:n-1]
	h.MinHeapify(0)
}
func (h *Heap) InsertMin(val int) {
	h.heap = append(h.heap, val)
	n := h.Size()
	curr := n - 1
	par := (curr - 1) / 2
	for par >= 0 && h.heap[par] > h.heap[curr] {
		swap(&h.heap[par], &h.heap[curr])
		curr = par
		par = (curr - 1) / 2
	}
}

// Decrease value of key at index ind to val
func (h *Heap) DecreaseKey(ind int, val int) {
	n := h.Size()
	if n <= ind {
		return
	}
	if ind == 0 {
		h.heap[ind] = val
	}
	par := (ind - 1) / 2
	for par >= 0 && h.heap[par] > h.heap[ind] {
		swap(&h.heap[par], &h.heap[ind])
		ind = par
		par = (ind - 1) / 2
	}
}

// Delete key at given index
func (h *Heap) DeleteMin(ind int) {
	n := h.Size()
	if n <= ind {
		return
	}
	h.DecreaseKey(ind, -999) // - INF
	h.ExtractMin()
}
