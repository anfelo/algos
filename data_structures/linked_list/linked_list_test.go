package linkedlist

import "testing"

func TestLinkedListInit(t *testing.T) {
	n1 := NewNode(1, nil)
	n2 := NewNode(2, n1)

	ll := NewLinkedList()

	if n2.Next.Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, n2.Next.Data)
	}

	if n1.Next != nil {
		t.Fatal("Expected next node to be nil")
	}

	if ll.Head != nil {
		t.Fatal("Expected head node to be nil")
	}
}

func TestLinkedListInsertFirst(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertFirst(1)

	if ll.Head == nil {
		t.Fatal("Expected head to be not nil")
	}

	if ll.Head.Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.Head.Data)
	}

	ll.InsertFirst(2)

	if ll.Head.Data != 2 {
		t.Fatalf("Expected %v but got %v", 2, ll.Head.Data)
	}

	if ll.Head.Next.Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.Head.Next.Data)
	}
}

func TestLinkedListSize(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertFirst(1)
	ll.InsertFirst(2)
	ll.InsertFirst(3)

	if ll.Size() != 3 {
		t.Fatalf("Expected %v but got %v", 3, ll.Size())
	}
}

func TestLinkedListGetFirst(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertFirst(1)

	if ll.GetFirst().Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.Size())
	}

	ll.InsertFirst(2)

	if ll.GetFirst().Data != 2 {
		t.Fatalf("Expected %v but got %v", 2, ll.Size())
	}
}

func TestLinkedListGetLast(t *testing.T) {
	ll := NewLinkedList()

	if ll.GetLast() != nil {
		t.Fatal("Expected node to be nil")
	}

	ll.InsertFirst(1)

	if ll.GetLast().Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.Size())
	}

	ll.InsertFirst(2)

	if ll.GetLast().Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.Size())
	}
}

func TestLinkedListClear(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertFirst(1)
	ll.InsertFirst(2)
	ll.InsertFirst(3)

	if ll.Size() != 3 {
		t.Fatalf("Expected %v but got %v", 3, ll.Size())
	}

	ll.Clear()

	if ll.Size() != 0 {
		t.Fatalf("Expected %v but got %v", 0, ll.Size())
	}
}

func TestLinkedListRemoveFirst(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertFirst(1)
	ll.InsertFirst(2)
	ll.InsertFirst(3)

	if ll.GetFirst().Data != 3 {
		t.Fatalf("Expected %v but got %v", 3, ll.GetFirst().Data)
	}

	ll.RemoveFirst()
	if ll.GetFirst().Data != 2 {
		t.Fatalf("Expected %v but got %v", 2, ll.GetFirst().Data)
	}

	ll.RemoveFirst()
	if ll.GetFirst().Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.GetFirst().Data)
	}

	ll.RemoveFirst()
	if ll.GetFirst() != nil {
		t.Fatal("Expected node to be nil")
	}
}

func TestLinkedListRemoveLast(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertFirst(1)
	ll.InsertFirst(2)
	ll.InsertFirst(3)

	if ll.GetLast().Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.GetLast().Data)
	}

	ll.RemoveLast()
	if ll.GetLast().Data != 2 {
		t.Fatalf("Expected %v but got %v", 2, ll.GetLast().Data)
	}

	ll.RemoveLast()
	if ll.GetLast().Data != 3 {
		t.Fatalf("Expected %v but got %v", 3, ll.GetLast().Data)
	}

	ll.RemoveLast()
	if ll.GetLast() != nil {
		t.Fatal("Expected node to be nil")
	}
}

func TestLinkedListInsertLast(t *testing.T) {
	ll := NewLinkedList()

	ll.InsertLast(1)
	if ll.GetLast().Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.GetFirst().Data)
	}

	ll.InsertLast(2)
	if ll.GetLast().Data != 2 {
		t.Fatalf("Expected %v but got %v", 2, ll.GetFirst().Data)
	}

	ll.InsertLast(3)
	if ll.GetLast().Data != 3 {
		t.Fatalf("Expected %v but got %v", 3, ll.GetFirst().Data)
	}
}
