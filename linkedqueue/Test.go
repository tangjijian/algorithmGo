package linkedqueue

type LinkQueue interface {
	Length() int
	EnQueue(value interface{})
	DeQueue() interface{}
}

type Node struct {
	data  interface{}
	pNext *Node
}
type LinkedQueue struct {
	front *Node
	rear  *Node
}

func NewLinkedQueue() *LinkedQueue {
	return &LinkedQueue{}
}

func (lq *LinkedQueue) EnQueue(value interface{}) {
	newNode := &Node{data: value, pNext: nil}
	if lq.front == nil {
		lq.front = newNode
		lq.rear = newNode
	} else {
		lq.rear.pNext = newNode
		lq.rear = newNode
	}

}
func (lq *LinkedQueue) DeQueue() interface{} {
	if lq.front == nil {
		return nil
	}
	data := lq.front.data
	if lq.front.pNext == nil { //只有一个节点
		lq.front = nil
		lq.rear = nil
		return data
	}
	lq.front = lq.front.pNext
	return data
}
