package levelwidth

type Node struct {
	Data     int
	Children []*Node
}

func (n *Node) Add(d int) {
	n.Children = append(n.Children, NewNode(d))
}

func (n *Node) Remove(d int) {
	cc := []*Node{}
	for _, v := range n.Children {
		if v.Data != d {
			cc = append(cc, v)
		}
	}
	n.Children = cc
}

func NewNode(d int) *Node {
	return &Node{Data: d}
}

type Tree struct {
	Root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) TraverseBF(f func(n *Node)) {
	arr := []*Node{}
	arr = append(arr, t.Root)
	for len(arr) != 0 {
		var curr *Node
		curr, arr = arr[0], arr[1:]
		arr = append(arr, curr.Children...)
		f(curr)
	}
}

func (t *Tree) TraverseDF(f func(n *Node)) {
	arr := []*Node{}
	arr = append(arr, t.Root)
	for len(arr) != 0 {
		var curr *Node
		curr, arr = arr[0], arr[1:]
		arr = append(curr.Children, arr...)
		f(curr)
	}
}
