package queueweave

import "testing"

func TestQueuePeekMethodToReturnLastItemOrFalse(t *testing.T) {
	expectedResult := []int{1, 2}
	q := NewQueue()

	if _, exists := q.Peek(); exists {
		t.Fatalf("Expected %t but got %t", false, exists)
	}

	q.Add(expectedResult[0])
	q.Add(expectedResult[1])

	if v, exists := q.Peek(); exists {
		if v != expectedResult[0] {
			t.Fatalf("Expected %v but got %v", expectedResult[1], v)
		}
	}

}

func TestWeaveQueueMethodToReturnNewQueue(t *testing.T) {
	q1 := NewQueue()
	q1.Add(1)
	q1.Add(2)
	q1.Add(3)

	q2 := NewQueue()
	q2.Add(4)
	q2.Add(5)
	q2.Add(6)
	q2.Add(7)

	expectedResult := []int{7, 6, 3, 5, 2, 4, 1}

	w := Weave(q1, q2)

	if len(w.Items) != len(expectedResult) {
		t.Fatalf("Expected %v but got %v", len(expectedResult), len(w.Items))
	}

	for i, v := range w.Items {
		if v != expectedResult[i] {
			t.Fatalf("Expected %v but got %v", expectedResult[i], v)
		}
	}

}
