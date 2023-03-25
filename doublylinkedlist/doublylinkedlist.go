package doublylinkedlist

import (
	"errors"
	"fmt"
)

type Node struct {
	Value    int
	next     *Node
	previous *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
	size int
}

func (d *DoublyLinkedList) Append(data int) {
	node := Node{Value: data}
	if d.head == nil {
		d.head = &node
		d.tail = d.head
	} else {
		node.previous = d.tail
		d.tail.next = &node
		d.tail = &node
	}
	d.size += 1
}

func (d *DoublyLinkedList) Prepend(data int) {
	node := Node{Value: data}
	if d.head == nil {
		d.head = &node
		d.tail = d.head
	} else {
		node.next = d.head
		d.head.previous = &node
		d.head = &node
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

func (d *DoublyLinkedList) traverseToIndex(index int) *Node {
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
	node := Node{Value: value}
	leader := d.traverseToIndex(index - 1)
	follower := leader.next
	leader.next = &node
	node.previous = leader
	node.next = follower
	follower.previous = &node
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
}

func (d *DoublyLinkedList) Search(data int) (int, error) {
	current := d.head
	var index int

	for current != nil {
		if current.Value == data {
			return index, nil
		}
		current = current.next
		index++
	}
	return -1, errors.New("not found")

}

func (d *DoublyLinkedList) Reverse() *DoublyLinkedList {
	if d.head.next == nil {
		return d
	}
	first := d.head
	second := first.next
	d.tail = d.head
	for second != nil {
		temp := second.next
		second.next = first
		first = second
		second = temp
	}
	d.head.next = nil
	d.head = first
	return d
}
