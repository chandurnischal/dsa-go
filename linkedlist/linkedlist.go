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

func (l *LinkedList) traverseToIndex(index int) *Node {
	current := l.head
	count := 0
	for count != index {
		current = current.next
		count += 1
	}
	return current
}

func (l *LinkedList) Insert(index, data int) {
	if index <= 0 {
		l.Prepend(data)
		return
	}
	if index >= l.size {
		l.Append(data)
		return
	}

	node := Node{value: data}
	leader := l.traverseToIndex(index - 1)
	trailer := leader.next
	leader.next = &node
	node.next = trailer
	l.size += 1
}

func (l *LinkedList) Remove(index int) {
	if index <= 0 {
		l.head = l.head.next
		l.size -= 1
		return
	}
	if index >= l.size {
		leader := l.traverseToIndex(l.size - 2)
		l.tail = leader
		leader.next = nil
		l.size -= 1
		return
	}
	leader := l.traverseToIndex(index - 1)
	unwanted := leader.next
	trailer := unwanted.next
	leader.next = trailer
	l.size -= 1
}
