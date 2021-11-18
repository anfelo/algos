package circularlist

func Circular(ll *LinkedList) bool {
	slow := ll.Head
	fast := ll.Head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}
