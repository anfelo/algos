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

func TestLinkedListGetAt(t *testing.T) {
	ll := NewLinkedList()

	if ll.GetAt(0) != nil {
		t.Fatal("Expected node to be nil")
	}

	ll.InsertFirst(1)

	if ll.GetAt(10) != nil {
		t.Fatal("Expected node to be nil")
	}

	ll.InsertFirst(2)
	ll.InsertFirst(3)

	if ll.GetAt(0).Data != 3 {
		t.Fatalf("Expected %v but got %v", 3, ll.GetAt(0).Data)
	}

	if ll.GetAt(1).Data != 2 {
		t.Fatalf("Expected %v but got %v", 2, ll.GetAt(1).Data)
	}

	if ll.GetAt(2).Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.GetAt(2).Data)
	}

	if ll.GetAt(3) != nil {
		t.Fatal("Expected node to be nil")
	}
}

func TestLinkedListRemoveAt(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertFirst(1)
	ll.InsertFirst(2)
	ll.InsertFirst(3)

	ll.RemoveAt(10)
	if ll.GetAt(0).Data != 3 {
		t.Fatalf("Expected %v but got %v", 3, ll.GetAt(0).Data)
	}

	ll.RemoveAt(1)

	if ll.GetAt(0).Data != 3 {
		t.Fatalf("Expected %v but got %v", 3, ll.GetAt(0).Data)
	}

	if ll.GetAt(1).Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.GetAt(1).Data)
	}

	if ll.GetAt(2) != nil {
		t.Fatal("Expected node to be nil")
	}

	ll.RemoveAt(0)

	if ll.GetAt(0).Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.GetAt(0).Data)
	}

	ll.RemoveAt(0)

	if ll.GetAt(0) != nil {
		t.Fatal("Expected node to be nil")
	}
}

func TestLinkedListInsertAt(t *testing.T) {
	ll := NewLinkedList()

	ll.InsertAt(1, 2)
	if ll.GetAt(0) == nil {
		t.Fatal("Expected node not to be nil")
	}
	if ll.GetAt(0).Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, ll.GetAt(0).Data)
	}

	ll.InsertAt(2, 1)
	if ll.GetAt(1).Data != 2 {
		t.Fatalf("Expected %v but got %v", 2, ll.GetAt(1).Data)
	}

	ll.InsertAt(3, 10)
	if ll.GetLast().Data != 3 {
		t.Fatalf("Expected %v but got %v", 3, ll.GetLast().Data)
	}
}

func TestLinkedListForEach(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertFirst(1)
	ll.InsertFirst(2)
	ll.InsertFirst(3)

	ll.ForEach(func(n *Node, i int) {
		n.Data += 10
	})

	if ll.GetAt(0).Data != 13 {
		t.Fatalf("Expected %v but got %v", 13, ll.GetAt(0).Data)
	}

	if ll.GetAt(1).Data != 12 {
		t.Fatalf("Expected %v but got %v", 12, ll.GetAt(1).Data)
	}

	if ll.GetAt(2).Data != 11 {
		t.Fatalf("Expected %v but got %v", 11, ll.GetAt(2).Data)
	}
}

func TestLinkedListGenerator(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertFirst(1)
	ll.InsertFirst(2)
	ll.InsertFirst(3)

	for v := range ll.Nodes() {
		v.Data += 10
	}

	if ll.GetAt(0).Data != 13 {
		t.Fatalf("Expected %v but got %v", 13, ll.GetAt(0).Data)
	}

	if ll.GetAt(1).Data != 12 {
		t.Fatalf("Expected %v but got %v", 12, ll.GetAt(1).Data)
	}

	if ll.GetAt(2).Data != 11 {
		t.Fatalf("Expected %v but got %v", 11, ll.GetAt(2).Data)
	}
}

func TestLinkedListIterator(t *testing.T) {
	ll := NewLinkedList()
	ll.InsertFirst(1)
	ll.InsertFirst(2)
	ll.InsertFirst(3)

	for it := NewLinkedListIterator(ll); it.MoveNext(); {
		it.Value().Data += 10
	}

	if ll.GetAt(0).Data != 13 {
		t.Fatalf("Expected %v but got %v", 13, ll.GetAt(0).Data)
	}

	if ll.GetAt(1).Data != 12 {
		t.Fatalf("Expected %v but got %v", 12, ll.GetAt(1).Data)
	}

	if ll.GetAt(2).Data != 11 {
		t.Fatalf("Expected %v but got %v", 11, ll.GetAt(2).Data)
	}
}
