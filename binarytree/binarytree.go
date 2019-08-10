package binarytree

//import "../queue"

type BinaryTree struct {
	root *Node
	size int
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

// TODO :complete this function
func (b *BinaryTree) Insert(val int) { //insert node in level order traversal fashion
	node := &Node{
		val: val,
	}
	if b.size == 0 {
		b.root = node
		b.size++
		return
	}

	//	q := &queue.Queue{}

}
