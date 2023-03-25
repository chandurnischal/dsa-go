package linkedlist

import (
	"errors"
	"fmt"
)

type node struct {
	Value int
	next  *node
}

type LinkedList struct {
	head *node
	tail *node
	size int
}

func (l *LinkedList) Append(data int) {
	newNode := node{Value: data}

	if l.head == nil {
		l.head = &newNode
		l.tail = l.head
	} else {
		l.tail.next = &newNode
		l.tail = &newNode
	}
	l.size++
}

func (l *LinkedList) Prepend(data int) {
	newNode := node{Value: data}

	if l.head == nil {
		l.head = &newNode
		l.tail = l.head
	} else {
		newNode.next = l.head
		l.head = &newNode
	}
	l.size++
}

func (l *LinkedList) traverseToIndex(index int) *node {
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

	newNode := node{Value: data}
	leader := l.traverseToIndex(index - 1)
	trailer := leader.next
	leader.next = &newNode
	newNode.next = trailer
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

func (l *LinkedList) Search(value int) (int, error) {
	current := l.head
	var index int
	for current != nil {
		if current.Value == value {
			return index, nil
		}
		current = current.next
		index += 1
	}
	return -1, errors.New("not found")
}

func (l *LinkedList) PrintList() {
	temp := l.head

	for temp != nil {
		if temp.next != nil {
			fmt.Printf("%d->", temp.Value)
		} else {
			fmt.Printf("%d", temp.Value)
		}
		temp = temp.next
	}
	fmt.Println()
}
