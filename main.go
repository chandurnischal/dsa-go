package main

import (
	"dsa/doublylinkedlist"
)

func main() {
	dll := doublylinkedlist.DoublyLinkedList{}
	l := []int{5, 4, 6, 3, 7, 2, 8, 1, 9}
	for _, v := range l {
		dll.Append(v)
	}
	dll.Print()
	dll.Reverse()
	dll.Print()
}
