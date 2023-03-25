package linkedlist

import "fmt"

type node struct {
	Value    int
	next     *node
	previous *node
}

type DoublyLinkedList struct {
	head *node
	tail *node
	size int
}

func (d *DoublyLinkedList) Append(data int) {
	newNode := node{Value: data}
	if d.head == nil {
		d.head = &newNode
		d.tail = d.head
	} else {
		newNode.previous = d.tail
		d.tail.next = &newNode
		d.tail = &newNode
	}
	d.size += 1
}

func (d *DoublyLinkedList) Prepend(data int) {
	newNode := node{Value: data}
	if d.head == nil {
		d.head = &newNode
		d.tail = d.head
	} else {
		newNode.next = d.head
		d.head.previous = &newNode
		d.head = &newNode
	}
	d.size++
}

func (d *DoublyLinkedList) Print() {
	temp := d.head

	for temp != nil {
		if temp.next != nil {
			fmt.Printf("%d<->", temp.Value)
		} else {
			fmt.Printf("%d", temp.Value)
		}
		temp = temp.next
	}
	fmt.Println()
}

func (d *DoublyLinkedList) traverseToIndex(index int) *node {
	current := d.head
	var count int

	for count != index {
		current = current.next
		count++
	}
	return current
}

func (d *DoublyLinkedList) Insert(index, value int) {
	if index <= 0 {
		d.Prepend(value)
		return
	}
	if index >= d.size {
		d.Append(value)
		return
	}
	newNode := node{Value: value}
	leader := d.traverseToIndex(index - 1)
	follower := leader.next
	leader.next = &newNode
	newNode.previous = leader
	newNode.next = follower
	follower.previous = &newNode
	d.size++
}

func (d *DoublyLinkedList) Remove(index int) {
	if index <= 0 {
		d.head = d.head.next
		d.head.previous = nil
		d.size--
		return
	}
	if index >= d.size {
		d.tail = d.tail.previous
		d.tail.next = nil
		d.size--
		return
	}
	leader := d.traverseToIndex(index - 1)
	unwanted := leader.next
	follower := unwanted.next
	leader.next = follower
	follower.previous = leader
	d.size--
	return
}
