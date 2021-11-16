package stack

import "testing"

func TestStackToPushPopAndPeek(t *testing.T) {
	s := NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)

	expectedResult := []int{1, 2, 3}

	if len(s.Items) != len(expectedResult) {
		t.Fatalf("Expected stack length %v but got %v", len(expectedResult), len(s.Items))
	}

	if v := s.Pop(); v != expectedResult[2] {
		t.Fatalf("Expected %v but got %v", expectedResult[2], v)
	}

	if v, exists := s.Peek(); exists {
		if v != expectedResult[1] {
			t.Fatalf("Expected %v but got %v", expectedResult[1], v)
		}
	}

}
