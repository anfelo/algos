package fromlast

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
	last := l.GetLast()
	if last == nil {
		l.Head = NewNode(d, nil)
		return
	}
	last.Next = NewNode(d, nil)
}

func (l *LinkedList) GetAt(index int) *Node {
	if l.Head == nil {
		return nil
	}
	i := 0
	curr := l.Head
	for i < index {
		if curr == nil {
			return curr
		}
		i++
		curr = curr.Next
	}
	return curr
}

func (l *LinkedList) RemoveAt(index int) {
	if l.Head == nil {
		return
	}

	if index == 0 {
		l.Head = l.Head.Next
		return
	}

	prev := l.GetAt(index - 1)
	if prev == nil || prev.Next == nil {
		return
	}

	prev.Next = prev.Next.Next
}

func (l *LinkedList) InsertAt(d int, index int) {
	n := l.GetAt(index - 1)
	if n == nil {
		l.InsertLast(d)
		return
	}
	n.Next = NewNode(d, n.Next)
}

func (l *LinkedList) ForEach(f func(n *Node, i int)) {
	count := 0
	curr := l.Head
	for curr != nil {
		f(curr, count)
		count++
		curr = curr.Next
	}
}

// Generator
func (l *LinkedList) Nodes() <-chan *Node {
	out := make(chan *Node)
	go func() {
		defer close(out)
		curr := l.Head
		for curr != nil {
			out <- curr
			curr = curr.Next
		}
	}()
	return out
}

// Iterator
type LinkedListIterator struct {
	list    *LinkedList
	current int
}

func NewLinkedListIterator(l *LinkedList) *LinkedListIterator {
	return &LinkedListIterator{l, -1}
}

func (li *LinkedListIterator) MoveNext() bool {
	li.current++
	return li.current < li.list.Size()
}

func (li *LinkedListIterator) Value() *Node {
	return li.list.GetAt(li.current)
}
