package linkedstack

type Node struct {
	pNext *Node
	value interface{}
}

func InitNode() *Node {
	n := new(Node)
	return n
}
func (n *Node) Push(data interface{}) {
	newNode := &Node{pNext: nil, value: data}
	ppnode := n.pNext
	n.pNext = newNode
	newNode.pNext = ppnode
}
func (n *Node) Pop() interface{} {
	if !n.IsEmpty() {
		data := n.pNext.value
		n.pNext = n.pNext.pNext
		return data
	}
	return nil
}
func (n *Node) IsEmpty() bool {
	if n.pNext == nil {
		return true
	} else {
		return false
	}
}
func (n *Node) Length() bool {
	length := 0
	for node := n.pNext; node.pNext != nil; {
		length++
	}
	return length == 0
}
