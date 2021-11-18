package circularlist

import "testing"

func TestCircularToReturnTruOrFalse(t *testing.T) {
	ll := NewLinkedList()
	n1 := NewNode(1, nil)
	n2 := NewNode(2, nil)
	n3 := NewNode(3, nil)
	ll.Head = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n2

	c := Circular(ll)

	if !c {
		t.Fatalf("Expected %t but got %t", true, c)
	}

	ll2 := NewLinkedList()
	ll2.InsertLast(1)
	ll2.InsertLast(2)
	ll2.InsertLast(3)

	c2 := Circular(ll2)

	if c2 {
		t.Fatalf("Expected %t but got %t", false, c2)
	}
}
