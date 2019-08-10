package binarytree

import "fmt"

func InOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	InOrderTraversal(node.left)
	fmt.Println(node.val)
	InOrderTraversal(node.right)
}

func PreOrderTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.val)
	PreOrderTraversal(node.left)
	PreOrderTraversal(node.right)
}

func PostOrderTraversal(node *Node) {
	if node == nil {
		return nil
	}
	PostOrderTraversal(node.right)
	PostOrderTraversal(node.left)
}
