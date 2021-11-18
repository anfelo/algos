package midpoint

import "testing"

func TestMidpointToReturnTheMidNodeOfLinkedList(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertLast(1)
	ll.InsertLast(2)
	ll.InsertLast(3)
	ll.InsertLast(4)
	ll.InsertLast(5)

	m := Midpoint(ll)

	if m.Data != 3 {
		t.Fatalf("Expected %v but got %v", 3, m.Data)
	}

	ll.InsertLast(6)

	if m.Data != 3 {
		t.Fatalf("Expected %v but got %v", 3, m.Data)
	}
}
