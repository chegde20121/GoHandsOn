package linkedlist

type LinkedList struct {
	Head *Node
	Tail *Node
}

type Node struct {
	Value int
	Next  *Node
}

func (ll *LinkedList) First() *Node {
	return ll.Head
}

func (ll *LinkedList) Last() *Node {
	return ll.Tail
}

// Push to insert data to List
func (ll *LinkedList) Push(value int) {
	node := &Node{Value: value}
	if ll.Head == nil {
		ll.Head = node
	} else {
		ll.Tail.Next = node
	}
	ll.Tail = node
}

func (node *Node) NextValue() *Node {
	return node.Next
}
