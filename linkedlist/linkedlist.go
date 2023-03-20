package linkedlist

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

func (l *LinkedList) Append(data int) {
	node := Node{value: data}

	if l.head == nil {
		l.head = &node
		l.tail = l.head
	} else {
		l.tail.next = &node
		l.tail = &node
	}
	l.size++
}

func (l *LinkedList) Prepend(data int) {
	node := Node{value: data}

	if l.head == nil {
		l.head = &node
		l.tail = l.head
	} else {
		node.next = l.head
		l.head = &node
	}
	l.size++
}

func (l *LinkedList) PrintList() {
	temp := l.head

	for temp != nil {
		if temp.next != nil {
			fmt.Printf("%d->", temp.value)
		} else {
			fmt.Printf("%d", temp.value)
		}
		temp = temp.next
	}
	fmt.Println()
}
