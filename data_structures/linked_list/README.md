# Definition

Ordered collection of data. The collection contains a number of nodes. Each node contains some data and a reference to the next node. Is like a chain of nodes.

- Head Node: First node
- Tail Node: Last node, does not have a reference to next node (null)

# Implementation

## Node Class

| Function | Arguments | Returns | Directions | Example |
| -- | -- | -- | -- | -- |
| constructor | (Data, Node) | Node | Creates a class instance to represent a node.  The node should have two properties, 'data' and 'next'. Accept both of these as arguments to the 'Node' constructor, then assign them to the instance as properties 'data' and 'next'. If 'next' is not provided to the constructor, then default its value to be 'null'. |
<pre>
	const n = new Node('There');
	n.data // 'Hi'
	n.next // null
	const n2 = new Node('Hi', n);
	n.next // returns n
</pre> |

## Linked List Class

| Function | Arguments | Returns | Directions | Example |
| -- | -- | -- | -- | -- |
| constructor | - | LinkedList | Create a class to represent a linked list. When created, a linked list should have *no* head node associated with it. The LinkedList instance will have one property, 'head', which is a reference to the first node of the linked list. By default 'head' should be 'null'. |
<pre>
	const list = new LinkedList();
	list.head // null
</pre> |
| insertFirst | (data) | - | Creates a new Node from argument 'data' and assigns the resulting node to the 'head' property. Make sure to handle the case in which the linked list already has a node assigned to the 'head' property. |
<pre>
	const list = new LinkedList();
	list.insertFirst('Hi There'); // List has one node
</pre> |
| size | - | (int) | Returns the number of nodes in the linked list. |
<pre>
	const list = new LinkedList();
	list.insertFirst('a');
	list.insertFirst('b');
	list.insertFirst('c');
	list.size(); // returns 3
</pre> |
| getFirst | - | (Node) | Returns the first node of the linked list. |
<pre>
	const list = new LinkedList();
	list.insertFirst('a');
	list.insertFirst('b');
	list.getFirst(); // returns Node instance with data 'b'
</pre> |
| getLast | - | (Node) | Returns the last node of the linked list. |
<pre>
	const list = new LinkedList();
	list.insertFirst('a');
	list.insertFirst('b');
	list.getLast(); // returns node with data 'a'
</pre> |
| clear | - | - | Empties the linked list of any nodes. |
<pre>
	const list = new LinkedList();
	list.insertFirst('a');
	list.insertFirst('b');
	list.clear();
	list.size(); // returns 0
</pre> |
| removeFirst | - | - | Removes only the first node of the linked list. The list's head should now be the second element. |
<pre>
	const list = new LinkedList();
	list.insertFirst('a');
	list.insertFirst('b');
	list.removeFirst();
	list.getFirst(); // returns node with data 'a'
</pre> |
| removeLast | - | - | Removes the last node of the chain. |
<pre>
	const list = new LinkedList();
	list.insertFirst('a');
	list.insertFirst('b');
	list.removeLast();
	list.size(); // returns 1
	list.getLast(); // returns node with data of 'b'
</pre> |
| insertLast | (Data) | - | Inserts a new node with provided data at the end of the chain. |
<pre>
	const list = new LinkedList();
	list.insertFirst('a');
	list.insertFirst('b');
	list.insertLast('c');
	list.getLast(); // returns node with data 'c'
</pre> |
| getAt | (int) | (Node) | Returns the node at the provided index. |
<pre>
	const list = new List();
	list.insertFirst('a');
	list.insertFirst('b');
	list.insertFirst('c');
	list.getAt(1); // returns node with data 'b'
</pre> |