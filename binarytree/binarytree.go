package binarytree

import "github.com/shweta1207/go/queue"

type BinaryTree struct {
	root *Node
	size int
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
	if b.root == nil {
		b.root = node
		b.size++
		return
	}
	totalnodes := b.size
	parentNode := (totalnodes + 1) / 2
	cnt := 0
	var top *Node
	q := &queue.Queue{}
	q.Push(b.root)
	for q.Size() > 0 {
		top = q.Front().(*Node)
		q.Pop()
		cnt++
		if cnt == parentNode {
			break
		}
		q.Push(top.left)
		q.Push(top.right)
	}
	b.size++
	if top.left == nil {
		top.left = node
		return
	}
	top.right = node
}
