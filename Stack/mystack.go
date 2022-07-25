package Stack

import "fmt"

type ArrayStack interface {
	Clear()
	Push(data interface{})
	Pop() interface{}
	IsFull() bool
	IsEmpty() bool
	Size() int
}

type myStack struct {
	cap        int
	currentPtr int
	dataSource []interface{}
}

func NewStack() ArrayStack {
	stack := new(myStack)
	stack.cap = 10
	stack.currentPtr = 0
	stack.dataSource = make([]interface{}, 0, 10)
	return stack
}
func (stack *myStack) Clear() {
	stack.dataSource = make([]interface{}, 0, 10)
}
func (stack *myStack) Push(data interface{}) {
	if !stack.IsFull() {
		stack.dataSource = append(stack.dataSource, data)
		stack.currentPtr++
	}
}
func (stack *myStack) Pop() interface{} {
	if !stack.IsEmpty() {
		stack.currentPtr--
		val := stack.dataSource[stack.currentPtr]
		stack.dataSource = stack.dataSource[:stack.currentPtr]
		fmt.Println(fmt.Sprint(stack.dataSource))
		return val
	} else {
		return nil
	}
}

func (stack *myStack) IsFull() bool {
	if stack.currentPtr == stack.cap {
		return true
	} else {
		return false
	}
}
func (stack *myStack) IsEmpty() bool {
	if stack.currentPtr == 0 {
		return true
	} else {
		return false
	}
}
func (stack *myStack) Size() int {
	return stack.currentPtr
}
