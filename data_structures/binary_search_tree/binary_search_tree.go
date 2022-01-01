package binarysearchtree

type Node struct {
	Data  int
	Left  *Node
	Right *Node
}

func NewNode(d int, l *Node, r *Node) *Node {
	return &Node{Data: d, Left: l, Right: r}
}

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

// Depth-First
func PreorderTraverse(n *Node, arr []int) []int {
	if n == nil {
		return arr
	}
	arr = append(arr, n.Data)
	arr = PreorderTraverse(n.Left, arr)
	arr = PreorderTraverse(n.Right, arr)
	return arr
}

// Depth-First
func InorderTraverse(n *Node, arr []int) []int {
	if n == nil {
		return arr
	}
	arr = PreorderTraverse(n.Left, arr)
	arr = append(arr, n.Data)
	arr = PreorderTraverse(n.Right, arr)
	return arr
}

// Depth-First
func PostorderTraverse(n *Node, arr []int) []int {
	if n == nil {
		return arr
	}
	arr = PreorderTraverse(n.Left, arr)
	arr = PreorderTraverse(n.Right, arr)
	arr = append(arr, n.Data)
	return arr
}
