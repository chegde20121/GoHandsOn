package doublylinked

type DoublyList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

func (dl *DoublyList) First() *Node {
	return dl.Head
}

func (dl *DoublyList) Last() *Node {
	return dl.Tail
}

func (dl *DoublyList) Push(val int) {
	node := &Node{Value: val}
	if dl.Head == nil {
		dl.Head = node
	} else {
		node.Prev = dl.Tail
		dl.Tail.Next = node
	}
	dl.Tail = node
}

func (dl *Node) NextVal() *Node {
	return dl.Next
}

func (dl *Node) Previous() *Node {
	return dl.Prev
}
