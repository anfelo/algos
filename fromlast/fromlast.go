package fromlast

func FromLast(ll *LinkedList, n int) *Node {
	slow := ll.GetFirst()
	fast := ll.GetFirst()

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow, fast = slow.Next, fast.Next
	}

	return slow
}
