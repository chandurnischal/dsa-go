package stack

import "errors"

type Node struct {
	Value int
	below *Node
}

type Stack struct {
	top    *Node
	bottom *Node
	size   int
}

func (s *Stack) Peek() (int, error) {
	if s.top == nil {
		return -1, errors.New("stack empty")
	}
	return s.top.Value, nil
}

func (s *Stack) Push(data int) {
	node := Node{Value: data}
	if s.size == 0 {
		s.top = &node
		s.bottom = &node
	} else {
		holdingPointer := s.top
		s.top = &node
		s.top.below = holdingPointer
	}
	s.size++
}

func (s *Stack) Pop() error {
	if s.size == 0 {
		return errors.New("stack empty")
	}
	s.top = s.top.below
	s.size--
	return nil
}
