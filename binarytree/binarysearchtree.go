package binarytree

func (b *BinaryTree) InsertBST(val int) {
	b.root = InsertNodeBST(b.root, val)
	b.size++

}

func InsertNodeBST(node *Node, val int) *Node {
	if node == nil {
		return &Node{
			val: val,
		}
	}
	if val < node.val {
		node.left = InsertNodeBST(node.left, val)
	} else {
		node.right = InsertNodeBST(node.right, val)
	}
	return node
}

func (b *BinaryTree) SearchBST(val int) bool {
	_, ok := SearchNodeBST(b.root, val)
	return ok
}

func SearchNodeBST(node *Node, val int) (*Node, bool) {
	if node == nil {
		return nil, false
	}
	if val == node.val {
		return node, true
	}
	if val < node.val {
		return SearchNodeBST(node.left, val)
	}
	return SearchNodeBST(node.right, val)
}

func DeleteNodeBST(node *Node, val int) *Node {
	if node == nil {
		return node
	}
	if val < node.val {
		node.left = DeleteNodeBST(node.left, val)
	} else if val > node.val {
		node.right = DeleteNodeBST(node.right, val)
	} else {
		// Delete current node
		// no child
		// if node.left==nil && node.right==nil {
		// 	return nil
		// }
		// has single child
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		minNode := FindMin(node.right)
		node.val = minNode
		node.right = DeleteNodeBST(node.right, minNode)
	}
	return node
}

func (b *BinaryTree) InorderSuccessor(val int) int {
	if b.root == nil {
		return 0
	}
	node, ok := SearchNodeBST(b.root, val)
	if !ok {
		return 0
	}
	return FindMin(node.right)
}

func (b *BinaryTree) InorderPredeccessor(val int) int {
	if b.root == nil {
		return 0
	}
	node, ok := SearchNodeBST(b.root, val)
	if !ok {
		return 0
	}
	return FindMax(node.left)
}

func FindMin(node *Node) int {
	if node == nil {
		return 0
	}
	for node.left != nil {
		node = node.left
	}
	return node.val
}

func FindMax(node *Node) int {
	if node == nil {
		return 0
	}
	for node.right != nil {
		node = node.right
	}
	return node.val
}
