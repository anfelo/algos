# Directions

- First implement a 'peek' method on the Queue. Peek should return the last element
  (The next to be removed) from the queue without removing it

- Implement the "weave" function. Weave receives two queues as arguments and combines the contents of each into a new, third queue.
  The third queue should contain the _alternating_ content of the two queues. the function should handle queues of different length without inserting null or undefined values into the new one.

**Do not** access the array inside of any queue, only use the "add", "remove" and "peek" functions.

# Example

```go
q1 := NewQueue()
q2 := NewQueue()

q1.Add(1)
q1.Add(2)

q2.Add(3)
q2.Add(4)

wq := Weave(q1, q2)
wq.Remove() // 1
wq.Remove() // 3
wq.Remove() // 2
wq.Remove() // 4
```

# Solution

## 1 Add Peek method

```go
type Queue struct {
	Items []int
}

func (q *Queue) Add(item int) {
	q.Items = append([]int{item}, q.Items...)
}

func (q *Queue) Remove() int {
	x, ss := q.Items[len(q.Items)-1], q.Items[:len(q.Items)-1]
	q.Items = ss
	return x
}

func (q *Queue) Peek() (int, bool) {
	if len(q.Items) > 0 {
		return q.Items[len(q.Items)-1], true
	}
	return 0, false
}

func NewQueue() *Queue {
	return &Queue{
		Items: []int{},
	}
}
```

## 2 Weave function

```go
func Weave(q1 *Queue, q2 *Queue) *Queue {
	w := NewQueue()
out:
	for {
		_, exists1 := q1.Peek()
		if exists1 {
			w.Add(q1.Remove())
		}

		_, exists2 := q2.Peek()
		if exists2 {
			w.Add(q2.Remove())
		}

		if !exists1 && !exists2 {
			break out
		}
	}
	return w
}
```
