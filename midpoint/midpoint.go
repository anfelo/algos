package midpoint

func Midpoint(ll *LinkedList) *Node {
	slow := ll.Head
	fast := ll.Head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}
