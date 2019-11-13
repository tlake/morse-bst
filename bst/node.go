package bst

type Node struct {
	Dot  *Node
	Dash *Node
	Char string
}

func NewNode(morse, char string) *Node {
	node := &Node{Char: char}
	return node
}
