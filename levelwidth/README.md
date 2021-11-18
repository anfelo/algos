# Directions

Given the root node of a tree, return an array where each element is the width of each level

# Example

```go
    1
   / \
  2   3
 / \   \
4  5    6

LevelWidth(root) // [1, 2, 3]
```

# Solution

## Pseudocode

* create a counters array
* create an array to store the nodes
* add the root node to the nodes array
* initialize the first item in the counters array with 0
* append to the end of the nodes array a stop symbol
* while the nodes array has length greater than 1
	* grab and remove the first item in nodes array
	* if the removed item is the stop symbol
		* append it to the end of the nodes array
		* initialize the next value in the counters array with 0
	* else
		* add 1 to the last item in counters array
		* append to the end of the nodes array all the children of the removed item
* return the counters array

```go
func LevelWidth(r *Node) []int {
	c := make([]int, 1)
	arr := []*Node{r, NewNode(-999)} // Stop symbol is -999
	for len(arr) > 1 {
		var curr *Node
		curr, arr = arr[0], arr[1:]
		if curr.Data == -999 {
			arr = append(arr, curr)
			c = append(c, 0)
		} else {
			c[len(c)-1]++
			arr = append(arr, curr.Children...)
		}
	}
	return c
}
```