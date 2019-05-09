package node

//Node struct with pointers to Left and Right and integer value is stored in num
type Node struct {
	Left  *Node
	Right *Node
	Num   int
}

//setter for left subtree
func (n Node) setLeft(node *Node) {
	n.Left = node
}

//setter for right subtree
func (n Node) setRight(node *Node) {
	n.Right = node
}

//setter for value
func (n Node) setNum(value int) {
	n.Num = value
}
