package handlers

import (
	"fmt"

	doublylinked "github.com/chegde20121/GoHandsOn/src/doublyLinked"
	linkedlist "github.com/chegde20121/GoHandsOn/src/linkedList"
)

func HandleLinkedList() {
	list := linkedlist.LinkedList{}
	list.Push(1)
	list.Push(2)
	list.Push(5)
	n := list.First()
	for {
		fmt.Println(n.Value)
		n = n.Next
		if n == nil {
			break
		}
	}
}

func HandleDoublyLinkedList() {
	list := doublylinked.DoublyList{}
	list.Push(1)
	list.Push(2)
	list.Push(5)
	n := list.Last()
	for {
		fmt.Println(n.Value)
		n = n.Prev
		if n == nil {
			break
		}
	}
}
