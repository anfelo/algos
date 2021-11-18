package fromlast

import "testing"

func TestFromLastToReturnNthNodeFromLast(t *testing.T) {
	l := NewLinkedList()
	l.InsertLast(1)
	l.InsertLast(2)
	l.InsertLast(3)
	l.InsertLast(4)

	if FromLast(l, 2).Data != 2 {
		t.Fatalf("Expected %v but got %v", 2, FromLast(l, 2))
	}
}
