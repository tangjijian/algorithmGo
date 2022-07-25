package linkedstack

type Node struct {
	pNext *Node
	value interface{}
}

func InitNode() *Node {
	n := new(Node)
	return n
}
func (n *Node) Push(data interface{}) *Node {
	newNode := &Node{pNext: nil, value: data}
	ppnode := n.pNext
	n.pNext = newNode
	newNode.pNext = ppnode
	return newNode
}
func (n *Node) Pop() {

	for node := n.pNext; node.pNext != nil; {

	}
}
func (n *Node) IsEmpty() bool {
	length := 0
	for node := n.pNext; node.pNext != nil; {
		length++
	}
	return length == 0
}
