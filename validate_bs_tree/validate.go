package validate

func Validate(n *Node, min interface{}, max interface{}) bool {
	if max != nil && n.Data > max.(int) {
		return false
	}

	if min != nil && n.Data < min.(int) {
		return false
	}

	if n.Left != nil && !Validate(n.Left, min, n.Data) {
		return false
	}

	if n.Right != nil && !Validate(n.Right, n.Data, max) {
		return false
	}

	return true
}
