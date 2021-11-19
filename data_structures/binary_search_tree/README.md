# Binary Search Tree

A variation of a regular tree that has a couple of constrains on the way the tree is structured.

* Every node can have max 2 child nodes, left node and right node
* Values to the left of the parent node have a lower value and values to the right have greater values than the parent nodes

The most common types of questions regarding binary search trees are inserting new nodes and validating that the tree is well structured

# Implementation

## Node

```go
type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(d int, l *Node, r *Node) *Node {
	return &Node{Data: d, Left: l, Right: r}
}
```

## Insert Method (Recursive)

```go
func (n *Node) InsertRecursive(d int) {
	// Base Cases
	if d < n.Data && n.Left == nil {
		n.Left = NewNode(d, nil, nil)
	} else if d > n.Data && n.Right == nil {
		n.Right = NewNode(d, nil, nil)
	} else if d < n.Data && n.Left != nil {
		// Recursive Cases
		n.Left.InsertRecursive(d)
	} else if d > n.Data && n.Right != nil {
		n.Right.InsertRecursive(d)
	}
}
```

## Insert Method (Depth-First)

**Not Recommended**: Ends up traversing more then needed in some cases

```go
func (n *Node) Insert(d int) {
	arr := []*Node{n}
	var curr *Node
	for len(arr) > 0 {
		curr, arr = arr[0], arr[1:]
		if d < curr.Data {
			// left
			if curr.Left != nil {
				arr = append(arr, curr.Left)
			} else {
				curr.Left = NewNode(d, nil, nil)
			}
		} else {
			// right
			if curr.Right != nil {
				arr = append(arr, curr.Right)
			} else {
				curr.Right = NewNode(d, nil, nil)
			}
		}
	}
}
```

## Contains Method (Recursive)

```go
func (n *Node) Contains(d int) *Node {
	// Base Cases
	if n.Data == d {
		return n
	}

	// Recursive Cases
	if d < n.Data && n.Left != nil {
		return n.Left.Contains(d)
	}

	if d > n.Data && n.Right != nil {
		return n.Right.Contains(d)
	}

	// Default Case
	return nil
}
```