package linkedlist

type Node struct {
	Data int
	Next *Node
}

func NewNode(d int, n *Node) *Node {
	return &Node{
		Data: d,
		Next: n,
	}
}

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (l *LinkedList) InsertFirst(d int) {
	l.Head = NewNode(d, l.Head)
}

func (l *LinkedList) Size() int {
	count := 0
	curr := l.Head
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count
}

func (l *LinkedList) GetFirst() *Node {
	return l.Head
}

func (l *LinkedList) GetLast() *Node {
	if l.Head == nil {
		return nil
	}
	curr := l.Head
	for curr.Next != nil {
		curr = curr.Next
	}
	return curr
}

func (l *LinkedList) Clear() {
	l.Head = nil
}

func (l *LinkedList) RemoveFirst() {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
}

func (l *LinkedList) RemoveLast() {
	if l.Head == nil {
		return
	}
	if l.Head.Next == nil {
		l.Head = nil
		return
	}
	prev := l.Head
	node := prev.Next
	for node.Next != nil {
		prev = node
		node = prev.Next
	}
	prev.Next = nil
}

func (l *LinkedList) InsertLast(d int) {
	// if l.Head == nil {
	// 	l.Head = NewNode(d, nil)
	// }
	// curr := l.Head
	// for curr.Next != nil {
	// 	curr = curr.Next
	// }
	// curr.Next = NewNode(d, nil)
	last := l.GetLast()
	if last == nil {
		l.Head = NewNode(d, nil)
		return
	}
	last.Next = NewNode(d, nil)
}
