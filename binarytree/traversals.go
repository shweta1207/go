package binarytree

import (
	"fmt"

	"github.com/shweta1207/go/queue"
)

func (b *BinaryTree) InOrder() {
	InOrderTraversal(b.root)
}

func (b *BinaryTree) PreOrder() {
	PreOrderTraversal(b.root)

}
func (b *BinaryTree) PostOrder() {
	PostOrderTraversal(b.root)

}

func (b *BinaryTree) LevelOrder() {
	LevelOrderTraversal(b.root)
}

func InOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	InOrderTraversal(node.left)
	fmt.Print(node.val, " ")
	InOrderTraversal(node.right)
}

func PreOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.val, " ")
	PreOrderTraversal(node.left)
	PreOrderTraversal(node.right)
}

func PostOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	PostOrderTraversal(node.right)
	PostOrderTraversal(node.left)
	fmt.Print(node.val, " ")
}

func LevelOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	q := queue.Queue{}
	q.Push(node)
	for q.Size() > 0 {
		top := q.Front().(*Node)
		q.Pop()
		fmt.Print(top.val, " ")
		if top.left != nil {
			q.Push(top.left)
		}
		if top.right != nil {
			q.Push(top.right)
		}
	}
}
