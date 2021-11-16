package queueweave

import "testing"

func TestQueuePeekMethodToReturnLastItemOrFalse(t *testing.T) {
	expectedResult := []int{1, 2}
	q := NewQueue()

	if _, exists := q.Peek(); exists {
		t.Fatalf("Expected %v but got %v", false, exists)
	}

	q.Add(expectedResult[0])
	q.Add(expectedResult[1])

	if v, exists := q.Peek(); exists {
		if v != expectedResult[0] {
			t.Fatalf("Expected %v but got %v", expectedResult[1], v)
		}
	}

}
