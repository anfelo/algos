# Directions

Given a linked list, return tru if the list is circular, false if is not.

# Examples

```go
l := NewLinkedList()
n1 := NewNode(1)
n2 := NewNode(2)
n3 := NewNode(3)
l.Head = n1
n1.Next = n2
n2.Next = n3
n3.Next = n2

Circular(l) // true
```

# Solution

## Pseudocode

* create a variable 'slow', assign the value of the linked list head
* create a variable 'fast', assign the value of the linked list head
* loop while the next 2 values of fast are not null
	* assign to slow the next value
	* assing to fast the value of the next 2 nodes
	* if slow is equal to fast; return true
* return false

```go
func Circular(ll *LinkedList) bool {
	slow := ll.Head
	fast := ll.Head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
```