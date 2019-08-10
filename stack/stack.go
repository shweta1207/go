package stack

type Stack struct {
	head *Node
	size int
}
type Node struct {
	val  interface{}
	next *Node
}

func (s *Stack) Push(val interface{}) {
	node := &Node{
		val: val,
	}
	s.size++
	if s.head == nil {
		s.head = node
		return
	}
	node.next = s.head
	s.head = node

}

func (s *Stack) Pop() {
	if s.head == nil {
		return
	}
	s.head = s.head.next
	s.size--
}

func (s *Stack) Top() interface{} {
	if s.head == nil {
		return 0
	}
	return s.head.val
}
func (s *Stack) Empty() bool {
	return s.size == 0
}
func (s *Stack) Size() int {
	return s.size
}
