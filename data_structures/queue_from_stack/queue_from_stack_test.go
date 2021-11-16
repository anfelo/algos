package queuefromstack

import (
	"testing"
)

func TestQueueFromStack(t *testing.T) {
	q := NewQueue()
	q.Add(1)
	q.Add(2)
	expectedResult := []int{2, 1}

	if v, exists := q.Peek(); exists {
		if v != expectedResult[1] {
			t.Fatalf("Expected %v but got %v", expectedResult[1], v)
		}
	} else {
		t.Fatalf("Expected %v but got %v", expectedResult[1], v)
	}

	v := q.Remove()
	if v != expectedResult[1] {
		t.Fatalf("Expected %v but got %v", expectedResult[1], v)
	}

}
