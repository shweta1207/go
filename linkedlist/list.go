package linkedlist

import "fmt"

type List struct {
	head   *Node
	length int
}

type Node struct {
	val  int
	next *Node
}

func (l *List) InsertNode(val int) { // insertion at beginning of list

	if l.length == 0 {
		l.head = &Node{
			val: val,
		}
		l.length = l.length + 1
		return
	}
	newNode := &Node{
		val:  val,
		next: l.head,
	}
	l.head = newNode
	l.length = l.length + 1

}

func (l *List) InsertAfterNode(prev *Node, val int) { // insertion after a given node
	if prev == nil {
		return
	}
	n := &Node{
		val:  val,
		next: prev.next,
	}
	prev.next = n
	l.length = l.length + 1
}

func (l *List) Length() int {
	return l.length
}

func (l *List) TraverseList() { // print all the elements
	curr := l.head
	for curr != nil {
		fmt.Println(curr.val)
		curr = curr.next
	}
}

func (l *List) FindValue(val int) bool { // find if a value is present
	curr := l.head
	for curr != nil {
		if curr.val == val {
			return true
		}
	}
	return false
}

// deleting a given key
func (l *List) DeleteNode(val int) {
	prev := &Node{}
	curr := l.head
	for curr != nil && curr.val != val {
		curr = curr.next
	}
	if curr == nil {
		return
	}
	prev.next = curr.next
}

// deleting node at a given position
func (l *List) DeleteAfterNode(prev *Node) {
	if prev == nil {
		return
	}
	prev.next = prev.next.next
}
