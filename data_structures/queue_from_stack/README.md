# Directions

Implement a Queue data structure using two stacks. **Do not** create an array of the "Queue" class. Queue should implement the "add", "remove", "peek" methods.

# Example

```go
q := NewQueue()
q.Add(1)
q.Add(2)
q.Peek() // returns 1
q.Remove() // return 1
q.Remove() // return 2
```

# Implementation

- create a queue class
- queue class has two stacks s1 and s2
- when add is called
	- call the push method on s1 with the new item
- when remove is called
	- while s1 has items (peek)
		- push to s2 the value of s1 pop
	- store value of s2 pop
	- while s2 has items (peek)
		- push to s1 the value of s2 pop
	return stored value
- when peek is called
	- while s1 has items (peek)
		- push to s2 the value of s1 pop
	- store value of s2 peek
	- while s2 has items (peek)
		- push to s1 the value of s2 pop
	return stored value

```go
type Queue struct {
	s1 *Stack
	s2 *Stack
}

func (q *Queue) Add(item int) {
	q.s1.Push(item)
}

func (q *Queue) Remove() int {
out1:
	for {
		if _, exists := q.s1.Peek(); exists {
			q.s2.Push(q.s1.Pop())
		} else {
			break out1
		}
	}

	r := q.s2.Pop()

out2:
	for {
		if _, exists := q.s2.Peek(); exists {
			q.s1.Push(q.s2.Pop())
		} else {
			break out2
		}
	}

	return r
}

func (q *Queue) Peek() (int, bool) {
out1:
	for {
		if _, exists := q.s1.Peek(); exists {
			q.s2.Push(q.s1.Pop())
		} else {
			break out1
		}
	}

	r, exists := q.s2.Peek()

out2:
	for {
		if _, exists := q.s2.Peek(); exists {
			q.s1.Push(q.s2.Pop())
		} else {
			break out2
		}
	}

	return r, exists
}

func NewQueue() *Queue {
	return &Queue{
		NewStack(),
		NewStack(),
	}
}

```