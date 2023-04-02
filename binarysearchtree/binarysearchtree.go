package binarysearchtree

type Node struct {
	Value int
	left  *Node
	right *Node
}

type BST struct {
	Root *Node
}

func (b *BST) Insert(data int) {
	node := Node{Value: data}
	if b.Root == nil {
		b.Root = &node
		return
	}
	current := b.Root
	for b.Root != nil {
		if data < current.Value {
			// move left
			if current.left == nil {
				current.left = &node
				return
			} else {
				current = current.left
			}
		} else {
			// move right
			if current.right == nil {
				current.right = &node
				return
			} else {
				current = current.right
			}
		}
	}
}

func (b *BST) Lookup(data int) *Node {
	if b.Root == nil {
		return nil
	}
	current := b.Root
	for current != nil {
		if data < current.Value {
			current = current.left
		} else if data > current.Value {
			current = current.right
		} else if data == current.Value {
			return current
		}
	}
	return nil
}

func (b *BST) BreadthFirstSearch() []int {
	current := b.Root
	list := []int{}
	queue := []*Node{current}
	for len(queue) > 0 {
		current = queue[0]
		queue = queue[1:]
		list = append(list, current.Value)
		if current.left != nil {
			queue = append(queue, current.left)
		}
		if current.right != nil {
			queue = append(queue, current.right)
		}
	}

	return list
}

func (b *BST) Inorder(node *Node) []int {
	res := []int{}
	if node != nil {
		res = append(res, b.Inorder(node.left)...)
		res = append(res, node.Value)
		res = append(res, b.Inorder(node.right)...)
	}

	return res
}

func (b *BST) Preoder(node *Node) []int {
	res := []int{}

	if node != nil {
		res = append(res, node.Value)
		res = append(res, b.Preoder(node.left)...)
		res = append(res, b.Preoder(node.right)...)
	}
	return res
}

func (b *BST) PostOrder(node *Node) []int {
	res := []int{}
	if node != nil {
		res = append(res, b.PostOrder(node.left)...)
		res = append(res, b.PostOrder(node.right)...)
		res = append(res, node.Value)
	}
	return res
}
