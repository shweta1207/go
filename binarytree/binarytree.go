package binarytree

import "../queue"

type BinaryTree struct {
	root *Node
}

type Node struct {
	val   int
	left  *Node
	right *Node
}

func (b *BinaryTree) Insert(val int) { //insert node in level order traversal fashion
	node := &Node{
		val: val,
	}
	q := &queue.Queue{}
}
