package tree

import "testing"

func TestNodeFactoryMethod(t *testing.T) {
	n := NewNode(1)

	if n.Data != 1 {
		t.Fatalf("Expected %v but got %v", 1, n.Data)
	}

	if len(n.Children) != 0 {
		t.Fatalf("Expected %v but got %v", 0, len(n.Children))
	}
}

func TestNodeAddMethod(t *testing.T) {
	n := NewNode(1)
	n.Add(2)

	if len(n.Children) != 1 {
		t.Fatalf("Expected %v but got %v", 0, len(n.Children))
	}

	if n.Children[0].Data != 2 {
		t.Fatalf("Expected %v but got %v", 2, n.Children[0].Data)
	}
}

func TestNodeRemoveMethod(t *testing.T) {
	n := NewNode(1)
	n.Add(2)
	n.Add(3)
	n.Add(4)

	if len(n.Children) != 3 {
		t.Fatalf("Expected %v but got %v", 3, len(n.Children))
	}

	n.Remove(3)

	if len(n.Children) != 2 {
		t.Fatalf("Expected %v but got %v", 2, len(n.Children))
	}
}

func TestTreeFactoryMethod(t *testing.T) {
	tr := NewTree()

	if tr.Root != nil {
		t.Fatal("Expected root to be nil")
	}
}

func TestTreeTraverseBreadthFirst(t *testing.T) {
	tr := NewTree()
	n1 := NewNode(1)
	n1.Add(2)
	n1.Add(3)
	n1.Children[0].Add(4)
	n1.Children[0].Add(5)
	tr.Root = n1

	expectedResult := []int{1, 2, 3, 4, 5}
	nn := []int{}
	tr.TraverseBF(func(n *Node) {
		nn = append(nn, n.Data)
	})

	if len(nn) != 5 {
		t.Fatalf("Expected %v but got %v", 5, len(nn))
	}

	for i, v := range nn {
		if v != expectedResult[i] {
			t.Fatalf("Expected %v but got %v", v, expectedResult[i])
		}
	}
}

func TestTreeTraverseDepthFirst(t *testing.T) {
	tr := NewTree()
	n1 := NewNode(1)
	n1.Add(2)
	n1.Add(3)
	n1.Children[0].Add(4)
	n1.Children[0].Add(5)
	tr.Root = n1

	expectedResult := []int{1, 2, 4, 5, 3}
	nn := []int{}
	tr.TraverseDF(func(n *Node) {
		nn = append(nn, n.Data)
	})

	if len(nn) != 5 {
		t.Fatalf("Expected %v but got %v", 5, len(nn))
	}

	for i, v := range nn {
		if v != expectedResult[i] {
			t.Fatalf("Expected %v but got %v", v, expectedResult[i])
		}
	}
}
