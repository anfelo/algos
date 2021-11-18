package queueweave

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
