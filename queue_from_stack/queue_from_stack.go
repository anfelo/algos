package queuefromstack

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
