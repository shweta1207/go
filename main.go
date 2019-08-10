package main

import (
	"fmt"
	"reflect"

	"./queue" //local usage
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
		val int
		name string
		Par *MyStruct
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
		val:5,
		name: "test",
	})
	fmt.Println(s.Size())
	vals,ok:=s.Top().(*MyStruct)
	fmt.Println(vals,ok)
	fmt.Println(vals.val,vals.name)
	s.Pop()
	fmt.Println(s.Size())



}
