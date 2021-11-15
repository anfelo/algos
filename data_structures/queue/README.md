# Definition

A queue is a container of items where each of the items is waiting to be processed.

The order in how the items are processed depends on the order of entrance to the queue. In this case, the queue processes the items in a First In First Out (FIFO) fashion. This means that items that enter the queue first are going to be processed first.

- Adding or Enqueuing
- Removing or Dequeuing

## Implementation in JavaScript

- Create an ES6 class
- Hold a reference to the items in an Array
- Expose only two methods, **unshift** (add) and **pop** (remove)

## Golang

´´´go
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

func NewQueue() *Queue {
	return &Queue{
		Items: []int{},
	}
}
´´´

## JavaScript

´´´javascript
class Queue {
    constructor() {
        this.items = [];
    }

    add(item) {
        this.items.unshift();
    }

    remove() {
        this.items.pop();
    }
}
´´´
