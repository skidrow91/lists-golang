package linkedlist

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Size int
	Head *Node
	Node *Node
}

func (list *LinkedList) Add(val int) int {
	node := new(Node)
	node.Data = val
	node.Next = nil

	if list.Head == nil {
		list.Head = node
		list.Node = node
	} else {
		last := list.Node
		for last.Next != nil {
			last = last.Next
		}
		last.Next = node
		list.Node = last
	}

	list.Size += 1

	return node.Data
}

func (list *LinkedList) Get(index int) int {
	last := list.Head
	i := 0
	for last.Next != nil {
		if (i >= index) {
			break
		}
		last = last.Next
		i++
	}
	return last.Data
}

