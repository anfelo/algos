# Directions

Return the "middle" node of a linked list. If the list has an even number of elements, return the node at the end of the first half of the list.
**Do not** use a counter variable, **do not** retrieve the size of the list, only iterate through the list one time.

# Example

```go
l := NewLinkedList()
l.InsertLast(1)
l.InsertLast(2)
l.InsertLast(3)
Midpoint(l) // returns {Data: 2}
```

# Solution

## Pseudocode

* create a variable 'slow', assign the value of the linked list head
* create a variable 'fast', assign the value of the linked list head
* loop while the next 2 values of fast are not null
	* assign to slow the next value
	* assing to fast the value of the next 2 nodes
* return slow


```go
func Midpoint(ll *LinkedList) *Node {
	slow := ll.Head
	fast := ll.Head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
```