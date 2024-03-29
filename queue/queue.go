package queue

type Queue struct {
	front *Node
	rear  *Node
	size  int
}

type Node struct {
	val  interface{}
	next *Node
}

func (q *Queue) Push(val interface{}) {
	node := &Node{
		val: val,
	}
	q.size++
	if q.front == nil {
		q.front = node
		q.rear = node
		return
	}
	q.rear.next = node
	q.rear = node
}

func (q *Queue) Pop() {
	if q.front == nil {
		return
	}
	if q.size == 1 {
		q.front = nil
		q.rear = nil
		q.size--
		return
	}
	q.front = q.front.next
	q.size--
}

func (q *Queue) Front() interface{} {
	if q.front == nil {
		return nil
	}
	return q.front.val
}
func (q *Queue) Rear() interface{} {
	if q.front == nil {
		return nil
	}
	return q.rear.val
}
func (q *Queue) Size() int {
	return q.size
}
func (q *Queue) Empty() bool {
	return q.size == 0
}
