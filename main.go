package main

import (
	"algorithm/ArrayList"
	"algorithm/Stack"
	"algorithm/circelqueue"
	"fmt"
	"strconv"
)

func main() {
	queue := circelqueue.InitQueue()
	circelqueue.EnQueue(queue, 1)
	circelqueue.EnQueue(queue, 2)
	circelqueue.EnQueue(queue, 3)
	circelqueue.EnQueue(queue, 4)
	circelqueue.EnQueue(queue, 5)

	fmt.Println(circelqueue.DeQueue(queue))
	fmt.Println(circelqueue.DeQueue(queue))
	fmt.Println(circelqueue.QueueLength(queue))
	fmt.Println(circelqueue.DeQueue(queue))
	fmt.Println(circelqueue.DeQueue(queue))
	fmt.Println(circelqueue.DeQueue(queue))
}
func main6() {
	arr := make([]int, 0, 10)
	//arr[1] = 4            // 不能赋值，len = 0时没有分配空间
	fmt.Println(arr[:10]) // 可以切片操作，生成的是一个新切片
}

func main5() {
	stack := Stack.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Size())
}
func main4() {

	list := ArrayList.NewArrayList()
	for i := 0; i < 5; i++ {
		list.Append("a" + strconv.Itoa(i))
	}
	//for i := 1; i < 25; i++ {
	//	list.Insert(i, "X5")
	//}
	//list.Delete(1)
	//fmt.Println(list.Size())

	for it := list.Iterator(); it.HasNext(); {
		fmt.Println(it.Next())
	}
	fmt.Println(list.String())
}
func main1() {
	arr := [...]int{1, 2, 3, 4, 5}
	s := arr[:5]
	s[2] = 100
	fmt.Println(fmt.Sprint(arr))
	fmt.Println(fmt.Sprint(s))
	fmt.Printf("len=%d,cap=%d", len(s), cap(s))
	fmt.Printf("len=%d,cap=%d", len(arr), cap(arr))
	//fmt.Println(arr[5])
}
func main2() {
	arr := make([]int, 0, 10)

	for i := 1; i < 6; i++ {
		arr = append(arr, i)
	}
	arr = arr[:6]
	fmt.Println(fmt.Sprint(arr))
	//fmt.Println(fmt.Sprint(s))

}
func main3() {
	arr := make([]int, 5, 10)
	arr[0] = 1
	fmt.Println(fmt.Sprint(arr))
}
