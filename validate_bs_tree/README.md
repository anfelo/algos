# Directions

Given a node, validate the binary search tree, ensuring that every node's left hand child is less than the parent node's value, and that every node's right hand child is greater than the parent.


# Solution (Recursion)

## Pseudocode

* start at the root node, max and min variables start null
* if max is not null and node value is greater than max
	* return false
* if min is not null and node value is lower than min
	* return false
* if left child is not null
	* if call to validate function with left child and assign max to the current node's value is false
		* return false
* else if the right child is not null
	* if call to validate function with left child and assign min to the current node's value is false
		* return false
* return true

```go
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
```