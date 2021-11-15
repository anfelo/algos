package queue

import "testing"

func TestQueueToAddAndRemoveItems(t *testing.T) {
	expectedResult := []int{1, 2}
	q := NewQueue()
	q.Add(expectedResult[0])
	q.Add(expectedResult[1])

	if len(q.Items) != len(expectedResult) {
		t.Fatalf("Expected %d but got %d", len(expectedResult), len(q.Items))
	}

	r := q.Remove()

	if len(q.Items) != 1 {
		t.Fatalf("Expected %d but got %d", 1, len(q.Items))
	}
	if r != expectedResult[0] {
		t.Fatalf("Expected %d but got %d", expectedResult[1], len(q.Items))
	}
}
