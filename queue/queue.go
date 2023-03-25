package queue

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
	after *Node
}

type Queue struct {
	front *Node
	back  *Node
	size  int
}

func (q *Queue) Peek() (int, error) {
	if q.size == 0 {
		return -1, errors.New("queue empty")
	}
	return q.front.Value, nil
}

func (q *Queue) Enqueue(data int) {
	node := Node{Value: data}
	if q.size == 0 {
		q.front = &node
		q.back = &node
	} else {
		q.back.after = &node
		q.back = &node
	}
	q.size++
}

func (q *Queue) Dequeue() error {
	if q.size == 0 {
		return errors.New("queue empty")
	}
	q.front = q.front.after
	q.size--
	return nil
}

func (q *Queue) Print() {
	temp := q.front

	for temp != nil {
		if temp.after != nil {
			fmt.Printf("%d<-", temp.Value)
		} else {
			fmt.Printf("%d", temp.Value)
		}
		temp = temp.after
	}
	fmt.Println()
}
