# Definition

Ordered collection of data. The collection contains a number of nodes. Each node contains some data and a reference to the next node. Is like a chain of nodes.

- Head Node: First node
- Tail Node: Last node, does not have a reference to next node (null)

# Implementation

## Node Class

| Function | Arguments | Returns | Directions | Example |
| -- | -- | -- | -- | -- |
| constructor | (Data, Node) | Node | Creates a class instance to represent a node.  The node should have two properties, 'data' and 'next'. Accept both of these as arguments to the 'Node' constructor, then assign them to the instance as properties 'data' and 'next'. If 'next' is not provided to the constructor, then default its value to be 'null'. | <pre>const n = new Node('There'); <br>n.data // 'Hi' <br>n.next // null<br>const n2 = new Node('Hi', n); <br>n.next // returns n</pre> |

## Linked List Class

| Function | Arguments | Returns | Directions | Example |
| -- | -- | -- | -- | -- |
| constructor | - | LinkedList | Create a class to represent a linked list. When created, a linked list should have *no* head node associated with it. The LinkedList instance will have one property, 'head', which is a reference to the first node of the linked list. By default 'head' should be 'null'. | <pre>const list = new LinkedList();<br>list.head // null</pre> |
| insertFirst | (data) | - | Creates a new Node from argument 'data' and assigns the resulting node to the 'head' property. Make sure to handle the case in which the linked list already has a node assigned to the 'head' property. | <pre>const list = new LinkedList();<br>list.insertFirst('Hi There'); // List has one node</pre> |
| size | - | (int) | Returns the number of nodes in the linked list. |<pre>const list = new LinkedList();<br>list.insertFirst('a');<br>list.insertFirst('b');<br>list.insertFirst('c');<br>list.size(); // returns 3</pre> |
| getFirst | - | (Node) | Returns the first node of the linked list. |<pre>const list = new LinkedList();<br>list.insertFirst('a');<br>list.insertFirst('b');<br>list.getFirst(); // returns Node instance with data 'b'</pre> |
| getLast | - | (Node) | Returns the last node of the linked list. |<pre>const list = new LinkedList();<br>list.insertFirst('a');<br>list.insertFirst('b');<br>list.getLast(); // returns node with data 'a'</pre> |
| clear | - | - | Empties the linked list of any nodes. |<pre>const list = new LinkedList();<br>list.insertFirst('a');<br>list.insertFirst('b');<br>list.clear();<br>list.size(); // returns 0</pre> |
| removeFirst | - | - | Removes only the first node of the linked list. The list's head should now be the second element. |<pre>const list = new LinkedList();<br>list.insertFirst('a');<br>list.insertFirst('b');<br>list.removeFirst();<br>list.getFirst(); // returns node with data 'a'</pre> |
| removeLast | - | - | Removes the last node of the chain. |<pre>const list = new LinkedList();<br>list.insertFirst('a');<br>list.insertFirst('b');<br>list.removeLast();<br>list.size(); // returns 1<br>list.getLast(); // returns node with data of 'b'</pre> |
| insertLast | (Data) | - | Inserts a new node with provided data at the end of the chain. |<pre><br>const list = new LinkedList();<br>list.insertFirst('a');<br>list.insertFirst('b');<br>list.insertLast('c');<br>list.getLast(); // returns node with data 'c'</pre> |
| getAt | (int) | (Node) | Returns the node at the provided index. |<pre>const list = new List();list.insertFirst('a');<br>list.insertFirst('b');<br>list.insertFirst('c');<br>list.getAt(1); // returns node with data 'b'</pre> |
| removeAt | (int) | - | Removes node at the provided index. |<pre>const list = new List();<br>list.insertFirst('a');<br>list.insertFirst('b');<br>list.removeAt(1);<br>list.getAt(1); // returns node with data 'a'</pre> |
| insertAt | (Data, integer) | - | Create an insert a new node at provided index. If index is out of bounds, add the node to the end of the list. |<pre>const list = new List();<br>list.insertFirst('a');<br>list.insertFirst('b');<br>list.insertFirst('c');<br>list.insertAt('Hi', 1)<br>list.getAt(1); // returns node with data 'Hi'</pre> |
| forEach | (function) | - | Calls the provided function with every node of the chain. | <pre>const list = new List();<br>list.insertLast(1);<br>list.insertLast(2);<br>list.insertLast(3);<br>list.insertLast(4);<br>list.forEach(node => {<br> node.data += 10;<br>});<br>list.getAt(0); // Returns node with data '11'</pre> |
| for range Loop | - | - | Linked list should be compatible as the subject of a 'for range' loop. | <pre>const list = new List();<br><br>list.insertLast(1);<br>list.insertLast(2);<br>list.insertLast(3);<br>list.insertLast(4);<br><br>for(let node of list) {<br>	node.data += 10;<br>}<br><br>node.getAt(1); // returns node with data 11</pre> |

## Golang

```go
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
```