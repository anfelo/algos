package queueweave

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
