package main

import (
	"fmt"
	"reflect"

	"./binarytree"
	"./queue" //local usage
	"./sorting"
	"./stack"
)

func main() {

	var a interface{} = 7
	fmt.Println(a)
	v := reflect.ValueOf(a)
	t := reflect.TypeOf(a)
	tt := reflect.TypeOf(t)
	fmt.Println(v, t, tt)
	//val,ok:=a.(int)
	val, ok := a.(int)
	fmt.Println(val, ok)
	q := queue.Queue{}
	s := stack.Stack{}
	q.Push(4)
	q.Push(5)
	q.Push(6)
	fmt.Println(q.Size())
	v1 := q.Front().(int)
	q.Pop()
	v2 := q.Front().(int)
	q.Pop()
	v3 := q.Front().(int)
	q.Pop()
	fmt.Println(v1, v2, v3)
	fmt.Println(q.Empty())
	q.Push("shweta")
	n, ok := q.Front().(string)
	fmt.Println(n, ok)
	fmt.Println(q.Front().(string))

	s.Push(4)
	s.Push(5)
	s.Push(6)
	fmt.Println(s.Size())
	val, ok = s.Top().(int)
	fmt.Println(val, ok)
	val, ok = s.Top().(int)
	fmt.Println(val, ok)
	s.Pop()
	val, ok = s.Top().(int)
	fmt.Println(val, ok)
	s.Pop()
	val, ok = s.Top().(int)
	fmt.Println(val, ok)
	s.Pop()
	fmt.Println(s.Empty(), s.Size())

	type MyStruct struct {
		val  int
		name string
		Par  *MyStruct
	}

	// s.Push(MyStruct{
	// 	val:5,
	// 	name: "test",
	// })
	// fmt.Println(s.Size())
	// vals,ok:=s.Top().(MyStruct)
	// fmt.Println(vals,ok)
	// fmt.Println(vals.val,vals.name)
	// s.Pop()
	// fmt.Println(s.Size())

	s.Push(&MyStruct{
		val:  5,
		name: "test",
	})
	fmt.Println(s.Size())
	vals, ok := s.Top().(*MyStruct)
	fmt.Println(vals, ok)
	fmt.Println(vals.val, vals.name)
	s.Pop()
	fmt.Println(s.Size())

	b := binarytree.BinaryTree{}
	b.Insert(2)
	b.Insert(3)
	b.Insert(4)
	b.Insert(5)
	b.Insert(6)
	b.LevelOrder()
	fmt.Println("")

	b.InOrder()
	fmt.Println("")

	b.PreOrder()
	fmt.Println("")

	b.PostOrder()
	fmt.Println("\n\n\n\n")

	bst := binarytree.BinaryTree{}
	bst.InsertBST(7)
	bst.InsertBST(9)
	bst.InsertBST(4)
	bst.InsertBST(5)
	bst.InsertBST(6)
	bst.LevelOrder()
	fmt.Println("")

	bst.InOrder()
	fmt.Println("")

	bst.PreOrder()
	fmt.Println("")

	bst.PostOrder()
	fmt.Println("")

	fmt.Println(bst.InorderSuccessor(7))
	fmt.Println(bst.InorderPredeccessor(7))

	arr := []int{6, 9, 34, 90, 3, 5, 1}
	sorting.SelectionSort(arr)
	fmt.Println(arr)
	arr = []int{4, 7, 9, 45, 3}
	sorting.HeapSort(arr)
	fmt.Println(arr)

}
