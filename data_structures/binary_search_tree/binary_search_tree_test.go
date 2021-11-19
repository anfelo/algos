package binarysearchtree

import (
	"testing"
)

func TestNodeFactoryMethod(t *testing.T) {
	l := NewNode(-1, nil, nil)
	r := NewNode(3, nil, nil)
	n := NewNode(1, l, r)

	if n.Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, n.Data)
	}

	if n.Left.Data != -1 {
		t.Fatalf("Expected %v but got %v", -1, n.Left.Data)

	}

	if n.Right.Data != 3 {
		t.Fatalf("Expected %v but got %v", 3, n.Right.Data)
	}
}

func TestNodeInsert(t *testing.T) {
	l := NewNode(-1, nil, nil)
	r := NewNode(3, nil, nil)
	n := NewNode(1, l, r)

	n.Insert(-2)

	if n.Left.Left == nil {
		t.Fatal("Expected node not to be nil")
	} else if n.Left.Left.Data != -2 {
		t.Fatalf("Expected %v but got %v", -2, n.Left.Left.Data)
	}

	n.Insert(0)

	if n.Left.Right == nil {
		t.Fatal("Expected node not to be nil")
	} else if n.Left.Right.Data != 0 {
		t.Fatalf("Expected %v but got %v", 0, n.Left.Right.Data)
	}

	n.Insert(10)

	if n.Right.Right == nil {
		t.Fatal("Expected node not to be nil")
	} else if n.Right.Right.Data != 10 {
		t.Fatalf("Expected %v but got %v", 10, n.Right.Right.Data)
	}

	n.Insert(11)

	if n.Right.Right.Right == nil {
		t.Fatal("Expected node not to be nil")
	} else if n.Right.Right.Right.Data != 11 {
		t.Fatalf("Expected %v but got %v", 11, n.Right.Right.Right.Data)
	}
}

func TestNodeInsertRecursive(t *testing.T) {
	l := NewNode(-1, nil, nil)
	r := NewNode(3, nil, nil)
	n := NewNode(1, l, r)

	n.InsertRecursive(-2)

	if n.Left.Left == nil {
		t.Fatal("Expected node not to be nil")
	} else if n.Left.Left.Data != -2 {
		t.Fatalf("Expected %v but got %v", -2, n.Left.Left.Data)
	}

	n.InsertRecursive(0)

	if n.Left.Right == nil {
		t.Fatal("Expected node not to be nil")
	} else if n.Left.Right.Data != 0 {
		t.Fatalf("Expected %v but got %v", 0, n.Left.Right.Data)
	}

	n.InsertRecursive(10)

	if n.Right.Right == nil {
		t.Fatal("Expected node not to be nil")
	} else if n.Right.Right.Data != 10 {
		t.Fatalf("Expected %v but got %v", 10, n.Right.Right.Data)
	}

	n.InsertRecursive(11)

	if n.Right.Right.Right == nil {
		t.Fatal("Expected node not to be nil")
	} else if n.Right.Right.Right.Data != 11 {
		t.Fatalf("Expected %v but got %v", 11, n.Right.Right.Right.Data)
	}
}

func TestNodeContains(t *testing.T) {
	l := NewNode(-1, nil, nil)
	r := NewNode(3, nil, nil)
	n := NewNode(1, l, r)
	n.Insert(-2)
	n.Insert(0)
	n.Insert(10)
	n.Insert(11)

	s1 := n.Contains(10)
	if s1 == nil {
		t.Fatal("Expected node not to be nil")
	} else if s1.Data != 10 {
		t.Fatalf("Expected %v but got %v", 10, s1.Data)
	}

	s2 := n.Contains(0)
	if s2 == nil {
		t.Fatal("Expected node not to be nil")
	} else if s2.Data != 0 {
		t.Fatalf("Expected %v but got %v", 0, s2.Data)
	}

	s3 := n.Contains(20)
	if s3 != nil {
		t.Fatal("Expected node to be nil")
	}
}
