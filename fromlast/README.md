# Directions

Given a linked list and integer n, return the element n spaces from the last node in the list. **Do not** call the "size" method of the linked list. Assume that n will always be less than the length of the list.

# Examples

```go
l := NewLinkedList()
l.InsertLast(1)
l.InsertLast(2)
l.InsertLast(3)
l.InsertLast(4)

FromLast(l, 2) // returns {Data: 2}
```

# Solution

## Pseudocode

* create a variable 'slow', assign the value of the linked list head
* create a variable 'fast', assign the value of the linked list head
* loop until n
	* advance fast to the next node
* loop while the next value of fast is not null
	* assign to slow the next value
	* assing to fast the next value
* return slow

```go
func FromLast(ll *LinkedList, n int) *Node {
	slow := ll.GetFirst()
	fast := ll.GetFirst()

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	return slow
}
```