# Definition

A Stack is a container of an ordered collection of records.

Items in a Stack are added to the top of the stack (Push) and the items are removed from the top as well (Pop). The order in how items are added and removed is called First In Last Out (FILO).

- Pushing items to the Stack (adding)
- Popping items from the Stack (removing)

# Implementation

- Create a class that has
- collection of items
- Push should append the item to the end of the stack
- Pop should remove an item from the end of the stack
- Peek should return the last item in the stack

```go
type Stack struct {
	Items []int
}

func (s *Stack) Push(item int) {
	s.Items = append(s.Items, item)
}

func (s *Stack) Pop() int {
	x, ss := s.Items[len(s.Items)-1], s.Items[:len(s.Items)-1]
	s.Items = ss
	return x
}

func (s *Stack) Peek() (int, bool) {
	if len(s.Items) > 0 {
		return s.Items[len(s.Items)-1], true
	}
	return 0, false
}

func NewStack() *Stack {
	return &Stack{
		Items: []int{},
	}
}
```
